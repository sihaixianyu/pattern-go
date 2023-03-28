package facade

import (
	"fmt"
	"log"
	"testing"
)

func TestFacade(t *testing.T) {
	fmt.Println()
	walletFacade := NewWalletFacade("abc", 1234)
	fmt.Println()

	if err := walletFacade.AddMoneyToWallet("abc", 1234, 10); err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	fmt.Println()
	if err := walletFacade.DeDuctMoneyFromWallet("abc", 1234, 5); err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}
