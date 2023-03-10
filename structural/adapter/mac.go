package adapter

import "fmt"

type Mac struct {}

func (m *Mac) InsertLightningPort() {
	fmt.Println("Lightning connector is pluged into mac machine")
}