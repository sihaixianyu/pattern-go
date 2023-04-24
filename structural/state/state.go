package state

import "fmt"

type state interface {
	addItem(int) error
	requestItem() error
	insertMoney(money int) error
	dispenseItem() error
}

type noItemState struct {
	vm *VendingMachine
}

func (s *noItemState) requestItem() error {
	return fmt.Errorf("Item out of stock")
}

func (s *noItemState) addItem(count int) error {
	s.vm.incrementItemCount(count)
	s.vm.setState(s.vm.hasItem)

	return nil
}

func (s *noItemState) insertMoney(money int) error {
	return fmt.Errorf("Item out of stock")
}

func (s *noItemState) dispenseItem() error {
	return fmt.Errorf("Item out of stock")
}

type hasItemState struct {
	vm *VendingMachine
}

func (s *hasItemState) requestItem() error {
	if s.vm.itemCount == 0 {
		s.vm.setState(s.vm.noItem)
		return fmt.Errorf("No item present")
	}

	fmt.Printf("Item requested\n")
	s.vm.setState(s.vm.itemRequested)
	return nil
}

func (s *hasItemState) addItem(count int) error {
	fmt.Printf("%d items added\n", count)
	s.vm.incrementItemCount(count)
	return nil
}

func (s *hasItemState) insertMoney(money int) error {
	return fmt.Errorf("Please select item first")
}

func (s *hasItemState) dispenseItem() error {
	return fmt.Errorf("Please select item first")
}

type ItemRequestState struct {
	vm *VendingMachine
}

func (s *ItemRequestState) requestItem() error {
	return fmt.Errorf("Item already requested")
}

func (s *ItemRequestState) addItem(count int) error {
	return fmt.Errorf("Item Dispense in progress")
}

func (s *ItemRequestState) insertMoney(money int) error {
	if money < s.vm.itemPrice {
		return fmt.Errorf("Inserted money is less. Please insert %d", s.vm.itemPrice)
	}

	fmt.Println("Money entered is enough")
	s.vm.setState(s.vm.hasMoney)
	return nil
}

func (s *ItemRequestState) dispenseItem() error {
	return fmt.Errorf("Please insert money first")
}

type hasMoneyState struct {
	vm *VendingMachine
}

func (s *hasMoneyState) requestItem() error {
	return fmt.Errorf("Item dispense in progress")
}

func (s *hasMoneyState) addItem(count int) error {
	return fmt.Errorf("Item dispense in progress")
}

func (s *hasMoneyState) insertMoney(money int) error {
	return fmt.Errorf("Item out of stock")
}

func (s *hasMoneyState) dispenseItem() error {
	fmt.Println("Dispense Item")
	s.vm.itemCount -= 1

	if s.vm.itemCount == 0 {
		s.vm.setState(s.vm.noItem)
	} else {
		s.vm.setState(s.vm.hasItem)
	}

	return nil
}
