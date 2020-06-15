// Add a function Withdraw(amount int) bool to the gopl.io/ch9/bank1 program. The result should
// indicate whether the transaction succeeded of failed due to insufficient funds. The message sent
// to the monitor goroutine must contain both the amount to withdraw and a new channel over which
// the monitor goroutine can send the boolean result back to Withdraw.
package bank

type withdrawal struct {
	amount int
	ok     chan<- bool
}

var deposits = make(chan int)            // send amount to deposit
var withdrawals = make(chan *withdrawal) // send withdrawal request
var balances = make(chan int)            // receive balance

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

func Withdraw(amount int) bool {
	ok := make(chan bool)
	withdrawals <- &withdrawal{amount, ok}
	return <-ok
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case w := <-withdrawals:
			ok := balance >= w.amount
			if ok {
				balance -= w.amount
			}
			w.ok <- ok
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
