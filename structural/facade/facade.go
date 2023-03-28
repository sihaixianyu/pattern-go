package facade

import "fmt"

type WalletFacade struct {
	account      Account
	wallet       Wallet
	securityCode SecurityCode
	notification Notification
	ledger       Ledger
}

func NewWalletFacade(account string, code int) WalletFacade {
	fmt.Println("Starting create account")

	var walletFacade = WalletFacade{
		account:      NewAccount(account),
		securityCode: NewSecurityCode(code),
		wallet:       NewWallet(),
		notification: Notification{},
		ledger:       Ledger{},
	}

	fmt.Println("Account created")

	return walletFacade
}

func (w *WalletFacade) AddMoneyToWallet(account string, securityCode int, amount int) error {
	fmt.Println("Starting add money to wallet")

	if err := w.account.CheckAccount(account); err != nil {
		return err
	}
	if err := w.securityCode.CheckCode(securityCode); err != nil {
		return err
	}

	w.wallet.CreditBalance(amount)
	w.notification.SendWalletCreditNotification()
	w.ledger.MakeEntry(account, "credit", amount)

	return nil
}

func (w *WalletFacade) DeDuctMoneyFromWallet(account string, securityCode int, amount int) error {
	fmt.Println("Starting debit money from wallet...")

	if err := w.account.CheckAccount(account); err != nil {
		return err
	}
	if err := w.securityCode.CheckCode(securityCode); err != nil {
		return err
	}
	if err := w.wallet.DebitBalance(amount); err != nil {
		return err
	}

	w.notification.SendWalletDebitNotification()
	w.ledger.MakeEntry(account, "debit", amount)

	return nil
}
