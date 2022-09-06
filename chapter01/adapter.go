package main

import (
	"fmt"

	"github.com/backmap/golang/DSA/chapter01/adapter"
)

// Iprocess interface
type Iprocess interface {
	process()
}

// Adapter struct
type Adapter struct {
	adaptee Adaptee
}

// Adapter class method: process
func (adapter Adapter) process() {
	fmt.Println(">>> Adapter process")
	adapter.adaptee.convert()
}

// Adaptee struct
type Adaptee struct {
	adapterType int
}

// Adaptee class method: convert
func (adaptee Adaptee) convert() {
	fmt.Println(">>> Adaptee Convert method --> adapterType:", adaptee.adapterType)
}

func testAdapter1() {
	fmt.Println("\n|-------------------------chapter01/adapter.go-------------------------|")
	fmt.Println(" - Define an Iprocess interface with a process() method")
	fmt.Println(" - Adaptee struct has a convert method.")
	fmt.Println(" - Adapter wraps the Adaptee. In its implemention of Iprocess interface,")
	fmt.Println(" Adapter invokes its adaptee's convert method.")
	fmt.Println("|----------------------------------------------------------------------|")
	fmt.Println()
	var processor Iprocess = Adapter{}
	processor.process()
}

func testAdapter2() {
	fmt.Println("\n|-----------------------------chapter01/adapter/*-----------------------------|")
	fmt.Println(" - A client interface \"Computer\" with an InsertIntoLightningPort() method.")
	fmt.Println(" - Client struct has a method which invokes the above method. ")
	fmt.Println(" - A Mac struct (Known Service) implements the Computer interface.")
	fmt.Println(" - A Windows struct (Unknown service) has an insertIntoUSBPort() method.")
	fmt.Println(" - To convert lightening signals to USB, we can define an WindowsAdapter struct")
	fmt.Println("  which wraps the Windows struct, and invokes its insertIntoUSBPort() method.")
	fmt.Println("|------------------------------------------------------------------------------|")
	fmt.Println()
	client := &adapter.Client{}
	mac := &adapter.Mac{}

	client.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := &adapter.Windows{}
	windowsMachineAdapter := adapter.NewWindowAdapter(windowsMachine)

	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}
