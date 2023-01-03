package enums

type Command byte

const (
	Haptic                = Command(iota) // 'H'
	Blink                                 // 'B'
	Shutdown                              // 'X'
	RequestHardwareStatus                 // 'T'
	DFUMode                               // 'D'
	PreventShutdown                       // 'Q'
	StartCalibration                      // 'C'
	AbortCalibration                      // 'A'
	SaveCalibration                       // 'S
)

type CommandPreventShutdownFirstArgument byte

const (
	On    = CommandPreventShutdownFirstArgument(iota) // "on"
	Off                                               // "off"
	NoUSB                                             // "usb-not-connected"
)

type CommandPreventShutdownSecondArgument byte

const (
	Pernament = CommandPreventShutdownSecondArgument(iota) // "perm"
	Temporary                                              // "temp"
)
