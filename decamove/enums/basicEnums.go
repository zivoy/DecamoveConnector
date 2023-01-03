package enums

type HapticEffect byte

const (
	SingleClickHaptic HapticEffect = 0x1
	DoubleClickHaptic HapticEffect = 0xA
	Buzz              HapticEffect = 0xE
	Calibration       HapticEffect = 0x18
)

type Accuracy byte

const (
	Unreliable = Accuracy(iota)
	Low
	Medium
	High
	UnknownAccuracy
)

type DecamoveState byte

const (
	EnteringSleep = DecamoveState(iota)
	LeavingSleep
	ShuttingDown
	SingleCLickResponse
	DoubleClickResponse
	TrupleClickResponse
	USBConnection
	USBDisconnection
)

type DongleState byte

const (
	Closed = DongleState(iota)
	Open
	Paired
	Streaming
)
