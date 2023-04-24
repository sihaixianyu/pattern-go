package state

import (
	"fmt"
	"log"
	"testing"
)

func TestState(t *testing.T) {
	vm := NewVendingMachine(1, 99)

	err := vm.RequestItem()
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = vm.InsertMoney(99)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vm.DispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vm.AddItem(2)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()

	err = vm.RequestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vm.InsertMoney(99)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vm.DispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
}
