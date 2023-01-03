package decamove

import (
	"fmt"
	"github.com/zivoy/decamoveConnector/decamove/enums"
)

func accMessage(accuracy enums.Accuracy) string {
	switch accuracy {
	case enums.Low:
		return "Low accuracy"
	case enums.High:
		return "high accuracy"
	case enums.Medium:
		return "Average accuracy"
	case enums.Unreliable:
		return "Accuracy is unreliable"
	}
	return "Unknown value"
}

func feedbackMessage(feedback enums.DecamoveState) string {
	switch feedback {
	case enums.SingleCLickResponse:
		return "Single click"
	case enums.DoubleClickResponse:
		return "Doubled clicked"
	case enums.TripleClickResponse:
		return "Triple clicked"
	case enums.EnteringSleep:
		return "Going to sleep"
	case enums.LeavingSleep:
		return "Waking up"
	case enums.ShuttingDown:
		return "Shutting down"
	case enums.USBConnection:
		return "Usb connected"
	case enums.USBDisconnection:
		return "USB disconnected"
	}
	return fmt.Sprintf("Unknown state %d", feedback)
}
