package adapter

import "fmt"

type Computer interface {
	InsertLightningPort()
}

type Mac struct{}

func (m *Mac) InsertLightningPort() {
	fmt.Println("Lightning connector is pluged into mac machine")
}

type Windows struct{}

func (w *Windows) insertIntoUSBPort() {
	fmt.Println("USB connector is plugged into windows machine.")
}
