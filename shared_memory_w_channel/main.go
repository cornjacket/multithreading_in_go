package main

import "time"

var (
	withdraw_done = false
	deposit_done = false
	channel = make(chan int)
)

func withdraw_thread() {
	for i := 0; i < 10000; i++ {
		transaction(-10)
	}
	println("Withdraw thread stop")
	withdraw_done = true
}

func deposit_thread() {
	for i := 0; i < 10000; i++ {
		transaction(10)
	}
	println("Deposti thread stop")
	deposit_done = true
}

func main() {
	go withdraw_thread()
	go deposit_thread()
	for ; !withdraw_done && !deposit_done; {
		time.Sleep(1000 * time.Millisecond)
	}
	print(balance())
	println()
}

func transaction(amount int) { channel<-amount}

func balance() int { return <-channel }

func teller() {
	var balance int = 100 // balance is confined to teller goroutine
	//var amount int
	for {
		select {
		case channel<-balance:
			print(balance)
			println()
		case amount := <-channel:
			balance+=amount
			print(balance)
			println()
		}
	}
}

func init() {
	go teller() // start the teller 
}
