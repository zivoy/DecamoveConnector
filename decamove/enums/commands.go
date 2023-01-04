package enums

type HapticEffect byte

const (
	SingleClickHaptic HapticEffect = 0x1
	DoubleClickHaptic HapticEffect = 0xA
	Buzz              HapticEffect = 0xE
	Calibration       HapticEffect = 0x18
)

type CommandPreventShutdownEnable byte

const (
	On    = CommandPreventShutdownEnable(iota) // "on"
	Off                                        // "off"
	NoUSB                                      // "usb-not-connected"
)

type CommandPreventShutdownTime byte

const (
	Pernament = CommandPreventShutdownTime(iota) // "perm"
	Temporary                                    // "temp"
)
