package decamove

import (
    "fmt"
    log "github.com/sirupsen/logrus"
    "github.com/zivoy/decamoveConnector/decamove/enums"
    "math"
    "time"
)

/*
commands:

Haptic                  H
Blink                   B
Shutdown                X
RequestHardwareStatus   T
DFUMode                 D
PreventShutdown         Q
StartCalibration        C
AbortCalibration        A
SaveCalibration         S

*/


func shutdown() string{
 return "X"
}

/* duration in seconds and frequcny in hz */
func blink(duration time.Duration, frequency int) string{
    if frequency < 1{
        frequency =1
    } else if frequency > 20{
        frequency =20
    }
    return fmt.Sprintf("B %d %d", int(math.Round(duration.Seconds())), frequency)
}

func haptic(effect enums.HapticEffect, clearQueue bool, repetitions int) string{
    if repetitions < 1{
        return stopHaptic()
    }
    if repetitions > 100{
        repetitions = 100
    }
    cmd := "H"
    if clearQueue{
        cmd = fmt.Sprintf("%s C",cmd)
    }
    return fmt.Sprintf("%s %d %d",cmd, repetitions, effect)
}

func stopHaptic() string{
    return "H C"
}

func startCalibratingIMU() string{
    return "C"
}

func abortCalibratingIMU() string{
    return "A"
}

func SaveIMUCalibration() string{
    return "S"
}

func requestHardwareStatus() string{
    return "T"
}

func enterDFUMode() string{
    return "D"
}

func disapleMoveShutdown(enable enums.CommandPreventShutdownEnable, permanent enums.CommandPreventShutdownTime) string{
    var enableFlag string
    switch enable{
    case enums.On:
        enableFlag = "on"
    case enums.Off:
        enableFlag = "off"
    case enums.NoUSB:
        enableFlag = "usb-not-connected"
    default:
        log.Warn("Invalid enable argument", enable)
        return ""
    }

    var permanentFlag string
    switch permanent{
    case enums.Pernament:
        permanentFlag = "perm"
    case enums.Temporary:
        permanentFlag = "temp"
    default:
        log.Warn("Invalid permanent argument", permanent)
        return ""
    }

    return fmt.Sprintf("Q shutdown %s %s",enableFlag, permanentFlag)
}
