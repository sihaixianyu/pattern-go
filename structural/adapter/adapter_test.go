package adapter

import "testing"

func TestAdpater(t *testing.T) {
	client := &Client{}
	mac := &Mac{}

	client.InsertLightningConnector(mac)

	windowsMachine := &Windows{}
	WindowsAdapter := &WindowsAdapter{
		windowsMachine: windowsMachine,
	}

	client.InsertLightningConnector(WindowsAdapter)
}
