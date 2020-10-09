package main

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
	for i := 0; i < 10000; i++ {
		lock.Lock()
		balance -= 10
		lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
	println("Withdraw thread stop")
	withdraw_done = true 
}

func deposit_thread() {
	for i := 0; i < 10000; i++ {
		lock.Lock()
		balance += 10 
		lock.Unlock()
		time.Sleep(1 * time.Millisecond)
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
	print(balance)
	println()
}
