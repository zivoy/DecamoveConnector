package decamove

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/zivoy/decamoveConnector/decamove/enums"
	"github.com/zivoy/decamoveConnector/decamove/types"
	"go.bug.st/serial"
	"strings"
)

type Decamove struct {
	connection serial.Port

	Battery             int
	Accuracy enums.Accuracy
	calibrationCounter  uint8
}

func Connect(port string) Decamove {
	dm := Decamove{}
	mode := &serial.Mode{
		BaudRate: 256000,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}
	conn, err := serial.Open(port, mode)

	if err != nil {
		log.Error(err)
		return dm
	}
	dm.connection = conn
	return dm
}

func (d Decamove) StartListner(ctx context.Context) {
	go func() {
		buff := make([]byte, 100)
		const dataMaxLength = 1024
		data := make([]byte, 0, dataMaxLength)

		for {
			select {
			case <-ctx.Done():
				log.Info("stop signal recived")
				return
			default:
				n, err := d.connection.Read(buff)
				if err != nil {
					log.Fatal(err)
				}
				if n == 0 {
					log.Debug("EOF reached, listener stopping")
					break
				}

				data = append(data, buff[:n]...)

				if strings.Contains(string(data), "\n") {
					data = []byte(strings.TrimSuffix(string(data), "\r\n"))
					d.Read(data)
					data = make([]byte, 0, dataMaxLength)
				} else if len(data) >= dataMaxLength {
					log.Error("Exceeded buffer length, reseting data")
					data = make([]byte, 0, dataMaxLength)
				}
			}
		}
	}()
}

//const responseSuffix = "\r\n"

func (d Decamove) Read(packet []byte) {
	command := parse(packet)

	switch command.GetCommand() {
	case enums.RotationUpdate:
		rotation := command.(types.RotationPacket)
        d.Accuracy =rotation.Accuracy
        log.Debug(rotation)
	case enums.CalibrationUpdate:
		if d.calibrationCounter++; d.calibrationCounter == 100 {
			d.calibrationCounter = 0
			d.Accuracy = command.(types.CalibrationPacket).Accuracy
            log.Info(accMessage(d.Accuracy))
		}
	case enums.Feedback:
		feedback := command.(types.FeedbackPacket)
        log.Info(feedbackMessage(feedback.Feedback))
	case enums.DeviceInfo:
		deviceInfo := command.(types.DeviceInfoPacket)
		log.Infof("dongle firmware: %s | dongle address: %s \t deacamove firmware: %s | decamove address: %s",
			deviceInfo.DongleVersion, deviceInfo.DongleAddress, deviceInfo.DecaMoveFirmware, deviceInfo.DecaMoveAddress)
	case enums.BatteryUpdate:
		battery := command.(types.BatteryPacket)
		d.Battery = int(battery.BatteryLevel)
		log.Infof("Battery level at %d%%", d.Battery)
	case enums.HardwareReponse:

	case enums.UnknownResponse:
		unknown := command.(types.UnknownPacket)
		log.WithFields(log.Fields{
			"raw bytes": unknown.Packet,
		}).Info("Unknown packet recived")
	default:
		log.WithFields(log.Fields{
			"command": command,
		}).Info("Unknown command recoived")
	}
}
