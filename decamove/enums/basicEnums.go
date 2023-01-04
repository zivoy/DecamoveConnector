package enums

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
	TripleClickResponse
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
