package decamove

import (
	log "github.com/sirupsen/logrus"
	"github.com/zivoy/decamoveConnector/decamove/enums"
)

type Decamove struct {
	Battery int
    CalibrationAccuracy uint8
    calibrationCounter uint8
}

func (d Decamove) Read(packet []byte) {
	command := parse(packet)
	switch command.getCommand() {
	case enums.RotationUpdate:
        message := command.(RotationPacket)
        log.Info(message)
	case enums.CalibrationUpdate:
        if d.calibrationCounter++; d.calibrationCounter==100{
            d.calibrationCounter=0
            d.CalibrationAccuracy = command.(CalibrationPacket).Accuracy
        }
	case enums.Feedback:
        feedback := command.(FeedbackPacket)
        log.Info(feedback.Feedback)

	case enums.UnknownResponse:
        unknown := command.(UnknownPacket)
        log.WithFields(log.Fields{
            "raw bytes": unknown.Packet,
            }).Debug("Unknown packet recived")
	default:
		log.WithFields(log.Fields{
			"command": command,
		}).Debug("Unknown command recoived")
	}
}
