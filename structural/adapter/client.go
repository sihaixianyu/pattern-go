package adapter

import (
	"fmt"
)

type Client struct{}

func (c *Client) InsertLightningConnector(com Computer) {
	fmt.Println("Client inserts Lightning into computer.")
	com.InsertLightningPort()
}
