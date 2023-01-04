package main

import (
	"context"
	"fmt"
	"github.com/zivoy/decamoveConnector/decamove"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
	//	"go.bug.st/serial/enumerator"
	"go.bug.st/serial"
)

func main() {

	// Retrieve the port list
	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		log.Fatal("No serial ports found!")
	}

	// Print the list of detected ports
	for _, port := range ports {
		fmt.Printf("Found port: %v\n", port)
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	dm := decamove.Connect(ports[0])
	dm.StartListner(ctx)

	<-signalChan
	cancel()
	time.Sleep(500 * time.Millisecond)
}

//func scan() {
//	ports, err := enumerator.GetDetailedPortsList()
//	if err != nil {
//		log.Fatal(err)
//	}
//	if len(ports) == 0 {
//		fmt.Println("No serial ports found!")
//		return
//	}
//	for _, port := range ports {
//		fmt.Printf("Found port: %s - %s\n", port.Name, port.Product)
//		if port.IsUSB {
//			fmt.Printf("   USB ID     %x:%x\n", port.VID, port.PID)
//			fmt.Printf("   USB serial %s\n", port.SerialNumber)
//		}
//	}
//}
