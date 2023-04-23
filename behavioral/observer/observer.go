package observer

import "fmt"

type Observer interface {
	update(string)
	getId() string
}

type Customer struct {
	Id string
}

func (c *Customer) update(itemName string) {
	fmt.Printf("Sending email to customer %s for item %s\n", c.Id, itemName)
}

func (c *Customer) getId() string {
	return c.Id
}
