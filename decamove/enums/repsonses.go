package enums

type Message byte

const (
	UnknownResponse        = Message(iota)
	RotationUpdate // xx
    CalibrationUpdate // mm
    Feedback // ff
    BatteryUpdate // bb
    DeviceInfo // vv
    HardwareReponse // tt
)
