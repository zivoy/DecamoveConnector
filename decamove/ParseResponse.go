package decamove

import (
	"encoding/binary"
	"github.com/zivoy/decamoveConnector/decamove/enums"
	"github.com/zivoy/decamoveConnector/decamove/types"
	"gonum.org/v1/gonum/num/quat"
	"math"
)

func parse(message []byte) types.MessagePacket {
	command := string(message[:2])
	packet := message[2:]
	packetLength := len(packet) // the length checks are porbably unneeded but im leaving them anyways
	if packetLength >= 8 && command == "xx" {
		num, acc := parseRotationUpdate(packet)
		return types.RotationPacket{Message: enums.RotationUpdate, Quaternion: num, Accuracy: acc}
	} else if packetLength == 8 && command == "mm" {
		return types.CalibrationPacket{Message: enums.CalibrationUpdate, Accuracy: enums.Accuracy(packet[6])}
	} else if packetLength == 1 && command == "ff" {
		return types.FeedbackPacket{Message: enums.Feedback, Feedback: enums.DecamoveState(packet[0])}
	} else if packetLength == 2 && command == "bb" {
		return types.BatteryPacket{Message: enums.BatteryUpdate, BatteryLevel: parseBatteryLevel(packet)}
	} else if packetLength == 24 && command == "vv" {
		return types.DeviceInfoPacket{Message: enums.DeviceInfo,
			DongleVersion:    parseFirmware(packet[:5]),
			DecaMoveFirmware: parseFirmware(packet[5:10]),
			DongleAddress:    parseBleAdress(packet[10:17]),
			DecaMoveAddress:  parseBleAdress(packet[17:]),
		}
	} else if packetLength >= 6 && command == "tt" {
		return generateHardwareSurveryResponse(packet)
	}

	return types.UnknownPacket{Message: enums.UnknownResponse, Packet: message}
}

const quatMagicConstant float32 = 6.103515625e-05

func parseRotationUpdate(message []byte) (quat.Number, enums.Accuracy) {
	accuracy := enums.UnknownAccuracy
	if len(message) > 8 {
		accuracy = enums.Accuracy(message[8])
	}

	// there is a check for if the dongle is streamign here in the original

	quaternian := make([]float64, 4)
	quatNumberBytes := make([]byte, 4)
	for i := 0; i < 4; i++ {
		quatNumberBytes[2] = message[i*2]
		quatNumberBytes[3] = message[i*2+1]
		bits := binary.LittleEndian.Uint32(quatNumberBytes)
		quaternian[i] = float64(math.Float32frombits(bits) * quatMagicConstant) // multipling by constant to unshrink the number
	}

	return quat.Number{Real: quaternian[0], Imag: quaternian[1], Jmag: quaternian[2], Kmag: quaternian[3]}, accuracy
}

func parseBatteryLevel(message []byte) int32 {
	num := float64(binary.LittleEndian.Uint16(message))
	return int32(math.Ceil(num / 10))
}

func parseBleAdress(message []byte) types.BLEAdress {
	address := types.BLEAdress{}
	address.IdPeer = message[0]&0b111 == 0b111
	address.Type = uint8(message[0]) & math.MaxInt8

	addr := make([]byte, 8)
	for i := 0; i < 6; i++ {
		addr[i] = message[i+1] // todo check later if have to put 0s at start instead
	}
	address.Address = binary.LittleEndian.Uint64(addr)

	return address
}

func parseFirmware(message []byte) types.FirmwareVersion {
	valid := message[0] != 0
	return types.FirmwareVersion{Valid: valid, Major: message[1], Minor: message[2], Patch: message[3], PreRelease: message[4]}
}

func generateHardwareSurveryResponse(message []byte) types.HardWareSurveyPacket {
	serveyData := types.HardWareSurveyPacket{Message: enums.HardwareReponse, RawData: message}

	data := message[0]
	serveyData.IsSleeping = /*        */ data&(1<<0) == (1 << 0)
	serveyData.UsbConnected = /*      */ data&(1<<1) == (1 << 1)
	serveyData.BQPGActive = /*        */ data&(1<<2) == (1 << 2)
	serveyData.Charging = /*          */ data&(1<<3) == (1 << 3)
	serveyData.CP2102Failure = /*     */ data&(1<<4) == (1 << 4)
	serveyData.IMUFailiure = /*       */ data&(1<<5) == (1 << 5)
	serveyData.BatteryFailiure = /*   */ data&(1<<6) == (1 << 6)
	serveyData.HapticsFailiure = /*   */ data&(1<<7) == (1 << 7)

	data = message[1]
	serveyData.HapticsInitFaliure = /**/ data&(1<<0) == (1 << 0)

	return serveyData
}
