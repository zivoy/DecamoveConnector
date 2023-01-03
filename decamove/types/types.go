package types

import (
	"fmt"
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