package enums

type Message byte

const (
	UnknownResponse        = Message(iota)
	RotationUpdate // xx
    CalibrationUpdate // mm
    Feedback
    BatteryUpdate
    DeviceInfo
    HardwareReponse
)
