package bank3

import "sync"

var (
	mu      sync.RWMutex // guards balance
	balance int
)

func Deposit(ammount int) {
	mu.Lock()
	balance += ammount
	mu.Unlock()
}

func Balance() int {
	mu.RLock()
	defer mu.RUnlock()
	return balance
}

func Withdraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()

	Deposit(-amount)
	if Balance() < 0 {
		Deposit(amount)
		return false // insufficient funds
	}

	return true
}
