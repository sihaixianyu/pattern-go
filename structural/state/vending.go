package state

import "fmt"

type VendingMachine struct {
	hasItem       state
	itemRequested state
	hasMoney      state
	noItem        state

	currentState state

	itemCount int
	itemPrice int
}

func NewVendingMachine(itemCount, itemPrice int) *VendingMachine {
	vm := &VendingMachine{
		itemCount: itemCount,
		itemPrice: itemPrice,
	}

	vm.hasItem = &hasItemState{vm: vm}
	vm.itemRequested = &ItemRequestState{vm: vm}
	vm.hasMoney = &hasMoneyState{vm: vm}
	vm.noItem = &noItemState{vm: vm}

	if itemCount > 0 {
		vm.setState(vm.hasItem)
	} else {
		vm.setState(vm.noItem)
	}

	return vm
}


func (vm *VendingMachine) RequestItem() error {
	return vm.currentState.requestItem()
}

func (vm *VendingMachine) AddItem(count int) error {
	return vm.currentState.addItem(count)
}

func (vm *VendingMachine) InsertMoney(money int) error {
	return vm.currentState.insertMoney(money)
}

func (vm *VendingMachine) DispenseItem() error {
	return vm.currentState.dispenseItem()
}

func (vm *VendingMachine) setState(s state) {
	vm.currentState = s
}

func (vm *VendingMachine) incrementItemCount(count int) {
	fmt.Printf("Adding %d items\n", count)
	vm.itemCount = vm.itemCount + count
}
