package adapter

import (
	"fmt"
)

type WindowsAdapter struct {
	windowsMachine *Windows
}

func (adapter *WindowsAdapter) InsertLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	adapter.windowsMachine.insertIntoUSBPort()
}
