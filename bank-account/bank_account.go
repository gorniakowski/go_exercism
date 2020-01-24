package account

import "sync"

//Account balance $$, status true if account open, mux to lock account
type Account struct {
	sync.Mutex
	balance int64
	status  bool
}

//Open opens accoutnt with initial deposit
func Open(initialDeposit int64) *Account {

	if initialDeposit < 0 {
		return nil
	}
	account := Account{balance: initialDeposit, status: true}
	return &account
}

//Close closes account and retruns payout
func (account *Account) Close() (payout int64, ok bool) {
	account.Lock()
	defer account.Unlock()
	if !account.status {
		return 0, false
	}
	account.status = false
	return account.balance, true

}

//Balance retruns cash :)
func (account *Account) Balance() (balance int64, ok bool) {
	account.Lock()
	defer account.Unlock()
	if !account.status {
		return 0, false
	}
	return account.balance, true
}

//Deposit allows desposit or withdrawl money from account
func (account *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	account.Lock()
	defer account.Unlock()
	newBalance = account.balance + amount
	if !account.status || newBalance < 0 {
		return 0, false
	}
	account.balance = newBalance
	return newBalance, true
}
