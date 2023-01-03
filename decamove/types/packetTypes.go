package types

import (
    "github.com/zivoy/decamoveConnector/decamove/enums"
    "gonum.org/v1/gonum/num/quat"
)

type MessagePacket interface {
    GetCommand() enums.Message
}
type RotationPacket struct {
    enums.Message
    Quaternion quat.Number
    Accuracy   enums.Accuracy
}

func (m RotationPacket) GetCommand() enums.Message { return m.Message }

type UnknownPacket struct {
    enums.Message
    Packet []byte
}

func (m UnknownPacket) GetCommand() enums.Message { return m.Message }

type CalibrationPacket struct {
    enums.Message
    Accuracy enums.Accuracy
}

func (m CalibrationPacket) GetCommand() enums.Message { return m.Message }

type FeedbackPacket struct {
    enums.Message
    Feedback enums.DecamoveState
}

func (m FeedbackPacket) GetCommand() enums.Message { return m.Message }

type BatteryPacket struct {
    enums.Message
    BatteryLevel int32
}

func (m BatteryPacket) GetCommand() enums.Message { return m.Message }

type DeviceInfoPacket struct {
    enums.Message
    DongleVersion    FirmwareVersion
    DongleAddress     BLEAdress
    DecaMoveFirmware FirmwareVersion
    DecaMoveAddress   BLEAdress
}

func (m DeviceInfoPacket) GetCommand() enums.Message { return m.Message }

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

func (m HardWareSurveyPacket) GetCommand() enums.Message { return m.Message }
