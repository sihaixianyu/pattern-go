package observer

import "testing"

func TestObserver(t *testing.T) {
	ps5 := NewItem("PS5")

	customer1 := &Customer{
		Id: "customer1@gmail.com",
	}
	customer2 := &Customer{
		Id: "customer2@gmail.com",
	}

	ps5.Register(customer1)
	ps5.Register(customer2)

	ps5.UpdateAvailability()
}
