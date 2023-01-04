package decamove

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.bug.st/serial"
)

func getSerialSettings() *serial.Mode {
	return &serial.Mode{
		BaudRate: 256000,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}
}

func startListner(ctx context.Context, move Decamove) {
	buff := make([]byte, 100)
	const dataMaxLength = 1024
	data := make([]byte, 0, dataMaxLength)

	defer func() {
		move.connection = nil
	}()

	for {
		select {
		case <-ctx.Done():
			log.Info("stop signal recived")
			return
		default:
			n, err := move.connection.Read(buff)
			if err != nil {
				log.Fatal(err)
			}
			if n == 0 {
				log.Debug("EOF reached, listener stopping")
				break
			}

			data = append(data, buff[:n]...)

			if data[len(data)-1] == '\n' { // maybe change it to scan till there is a newline and cut in case it concatinates two messages
				move.Read(data)
				data = make([]byte, 0, dataMaxLength)
			} else if len(data) >= dataMaxLength {
				log.Warn("Exceeded buffer length, reseting data")
				data = make([]byte, 0, dataMaxLength)
			}
		}
	}
}

func write(d Decamove, command string) error {
	command = command + "\n"
	_, err := d.connection.Write([]byte(command))
	return err
}
