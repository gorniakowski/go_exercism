package account

import "sync"
//Account balance $$, status true if account open, mux to lock account 
type Account struct {
	balance int64
	status  bool
	mux     sync.Mutex
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
	account.mux.Lock()
	if !account.status {
		payout,ok = 0, false
	}else{
		payout,ok = account.balance,true
		account.status = false
	}		
	defer account.mux.Unlock()
	return payout, ok

}
//Balance retruns cash :)
func (account *Account) Balance() (balance int64, ok bool) {

	if !account.status {
		return 0, false
	}
	return account.balance, true
}
//Deposit allows desposit or withdrawl money from account
func (account *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	account.mux.Lock()
	newBalance = account.balance + amount
	if !account.status || newBalance < 0 {
		newBalance,ok = 0,false
	}else{
		account.balance,ok = newBalance, true
	}
	defer account.mux.Unlock()
	return newBalance, ok
}
