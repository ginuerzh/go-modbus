package main

import (
	"fmt"
	"time"

	"go.bug.st/serial"

	modbus "github.com/ginuerzh/go-modbus"
)

func main() {
	p := modbus.NewRTUClientProvider(modbus.WithEnableLogger(),
		modbus.WithSerialConfig("COM11", serial.Mode{
			BaudRate: 115200,
			DataBits: 8,
			StopBits: serial.OneStopBit,
			Parity:   serial.NoParity,
		}))

	client := modbus.NewClient(p)
	err := client.Connect()
	if err != nil {
		fmt.Println("connect failed, ", err)
		return
	}
	defer client.Close()

	fmt.Println("starting")
	for {
		results, err := client.ReadHoldingRegisters(1, 0x4001, 5)
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Printf("ReadHoldingRegisters %v\n", results)

		time.Sleep(time.Second * 1)
	}
}
