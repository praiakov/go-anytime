// o pacote bank oferece um banco seguro para concorrência, com uma conta
package bank

var deposits = make(chan int) // envia um valor para depósito
var balances = make(chan int) // recebe o saldo

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

func teller() {
	var balance int // balance está confinada na gorrotina teller
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // inici a gorrotina monitora
}
