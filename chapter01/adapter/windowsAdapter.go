package adapter

import "fmt"

type WindowsAdapter struct {
	windowMachine *Windows
}

func NewWindowAdapter(windowMachine *Windows) *WindowsAdapter {
	return &WindowsAdapter{windowMachine: windowMachine}
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println(">>> Adapter converts Lightning signal to USB.")
	w.windowMachine.insertIntoUSBPort()
}
