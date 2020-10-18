package main

// The purpose of this example is to show what would happen if the withdraw_thread
// removes more ($20) than the deposit_thread ($10) and so the balance goes below
// zero.

import (
	"sync"
	"time"
)

var (
	balance = 100
	lock = sync.Mutex{}
	withdraw_done = false
	deposit_done = false
)

func withdraw_thread() {
	println("Withdraw thread start")
	for i := 0; i < 1000; i++ {
		lock.Lock()
		balance -= 20
		lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
	println("Withdraw thread stop")
	withdraw_done = true 
}

func deposit_thread() {
	println("Deposit thread start")
	for i := 0; i < 1000; i++ {
		lock.Lock()
		balance += 10 
		lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
	println("Deposit thread stop")
	deposit_done = true
}

func main() {
	go withdraw_thread()
	go deposit_thread()
	for ; !withdraw_done || !deposit_done; {
		time.Sleep(1000 * time.Millisecond)
	}
	print(balance)
	println()
}
