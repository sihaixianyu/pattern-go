package adapter

import (
	"fmt"
	"testing"
)

type Client struct{}

func (c *Client) InsertLightningConnector(com Computer) {
	fmt.Println("Client inserts Lightning into computer.")
	com.InsertLightningPort()
}

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
