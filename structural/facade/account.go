package facade

import "fmt"

type Account struct {
	name string
}

func NewAccount(name string) Account {
	return Account{
		name: name,
	}
}

func (a *Account) CheckAccount(name string) error {
	if	a.name != name {
		return fmt.Errorf("Account name is incorrect!")
	}

	fmt.Println("Account is verified!")

	return nil
}
