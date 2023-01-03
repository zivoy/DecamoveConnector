package decamove

import (
	"fmt"
	"github.com/zivoy/decamoveConnector/decamove/enums"
	"gonum.org/v1/gonum/num/quat"
	"math"
)

type BLEAdress struct {
	IdPeer  bool
	Type    byte
	Address uint64
}

func (b BLEAdress) String() string {
	if b.Address == 0 {
		return "<invalid>"
	}
	return fmt.Sprint(b.Address)
}

type FirmwareVersion struct {
	Valid      bool
	Major      byte
	Minor      byte
	Patch      byte
	PreRelease byte
}

func (f FirmwareVersion) String() string {
	str := fmt.Sprintf("%d.%d.%d", f.Major, f.Minor, f.Patch)
	if f.PreRelease != math.MaxUint8 {
		return fmt.Sprintf("%s-pre%d", str, f.PreRelease)
	}
	return str
}

type MessagePacket interface {
	getCommand() enums.Message
}
type RotationPacket struct {
	enums.Message
	Quaternion quat.Number
	Accuracy   enums.Accuracy
}

func (m RotationPacket) getCommand() enums.Message { return m.Message }

type UnknownPacket struct {
	enums.Message
	Packet []byte
}

func (m UnknownPacket) getCommand() enums.Message { return m.Message }

type CalibrationPacket struct {
	enums.Message
	Accuracy uint8
}

func (m CalibrationPacket) getCommand() enums.Message { return m.Message }

type FeedbackPacket struct {
	enums.Message
	Feedback enums.DecamoveState
}

func (m FeedbackPacket) getCommand() enums.Message { return m.Message }

type BatteryPacket struct {
	enums.Message
	BatteryLevel int32
}

func (m BatteryPacket) getCommand() enums.Message { return m.Message }

type DeviceInfoPacket struct {
	enums.Message
	DongleVersion    FirmwareVersion
	DecaMoveFirmware FirmwareVersion
	DongleAdress     BLEAdress
	DecaMoveAdress   BLEAdress
}

func (m DeviceInfoPacket) getCommand() enums.Message { return m.Message }

type HardWareSurveyPacket struct {
	enums.Message
	RawData []byte

	IsSleeping         bool
	UsbConnected       bool
	BQPGActive         bool
	Charging           bool
	CP2102Failure      bool
	IMUFailiure        bool
	BatteryFailiure    bool
	HapticsFailiure    bool
	HapticsInitFaliure bool
}

func (m HardWareSurveyPacket) getCommand() enums.Message { return m.Message }
