package main

// The purpose of this example is to show how to use a condition varialbe to work with a process that
// is waiting on a specific condition. For this example that will prevent the balance from going below
// zero.
// The deposit_thread deposits $10 for 1000 deposits, while the withdraw_thread withdraws $20 for 500 withdraws.
// Note that the balance never goes below zero, unlike the non-condition variable version.

import (
	"sync"
	"time"
)

var (
	balance = 100
	lock = sync.Mutex{}
	moneyAvailable = sync.NewCond(&lock)
	withdraw_done = false
	deposit_done = false
)

func withdraw_thread() {
	println("Withdraw thread start")
	for i := 0; i < 500; i++ { // reduce number of withdraws to 500 to compensate for twice as much withdraw amount
		lock.Lock()
		for balance - 20 < 0 {
			moneyAvailable.Wait() // lock will be given up if balalnce < 20
		}
		balance -= 20
		println("Withdraw thread: ", balance)
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
		println("Deposit thread: ", balance)
		moneyAvailable.Signal() // signal the process that is waiting on moneyAvailable
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
	println("Main stop: ", balance)
}
