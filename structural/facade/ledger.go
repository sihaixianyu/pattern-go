package facade

import "fmt"

type Ledger struct{}

func (s *Ledger) MakeEntry(account, txnType string, amount int) {
	fmt.Printf("Make ledger entry for account %s with txnType %s for amount %d\n ", account, txnType, amount)
}
