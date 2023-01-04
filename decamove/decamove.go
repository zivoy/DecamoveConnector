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

	Battery            int
	Accuracy           enums.Accuracy
	calibrationCounter uint8

	DecamoveVersion types.FirmwareVersion
	DecamoveAddress types.BLEAdress

	DongleVersion types.FirmwareVersion
	DongleAddress types.BLEAdress
}

func Connect(port string) Decamove {
	dm := Decamove{}
	conn, err := serial.Open(port, getSerialSettings())
	if err != nil {
		log.Error(err)
		return dm
	}
	err = conn.SetDTR(true)
	if err != nil {
		log.Error(err)
		return dm
	}

	dm.connection = conn
	return dm
}

func (d Decamove) StartListner(ctx context.Context) {
	go startListner(ctx, d)
}

const responseSuffix = "\r\n"

func (d Decamove) Read(packet []byte) {
	packet = []byte(strings.TrimRight(string(packet), responseSuffix))
	command := parse(packet)

	switch command.GetCommand() {
	case enums.RotationUpdate:
		rotation := command.(types.RotationPacket)
		d.Accuracy = rotation.Accuracy
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
		d.DongleVersion = deviceInfo.DongleVersion
		d.DongleAddress = deviceInfo.DongleAddress
		d.DecamoveVersion = deviceInfo.DecaMoveFirmware
		d.DecamoveAddress = deviceInfo.DecaMoveAddress
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

func (d Decamove) DongleState() enums.DongleState {
	if d.connection == nil {
		return enums.Closed
	}
	// todo
	return 0
}
