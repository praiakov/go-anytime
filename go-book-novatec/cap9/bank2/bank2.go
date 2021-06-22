package bank2

var (
	sema    = make(chan struct{}, 1) // um semáforo binário que protege balance
	balance int
)

func Deposit(amount int) {
	sema <- struct{}{} // adquire o token
	balance = balance + amount
	<-sema // libera token
}

func Balance() int {
	sema <- struct{}{} // adquire o token
	b := balance
	<-sema // libera o token
	return b
}
