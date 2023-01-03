package enums

type CommandPreventShutdownEnable byte

const (
    On    = CommandPreventShutdownEnable(iota) // "on"
	Off                                               // "off"
	NoUSB                                             // "usb-not-connected"
)

type CommandPreventShutdownTime byte

const (
    Pernament = CommandPreventShutdownTime(iota) // "perm"
	Temporary                                              // "temp"
)
