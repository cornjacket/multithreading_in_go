package main

// This program should run along for a short time and then get into a deadlock situation.

import (
	"fmt"
	"sync"
	"time"
)

var (
	lock1 = sync.Mutex{}
	lock2 = sync.Mutex{}
)

func process1() {
	for {
		fmt.Println("Process 1 aquiring Lock 1\n")
		lock1.Lock()
		fmt.Println("Process 1 aquiring Lock 2\n")
		lock2.Lock()
		fmt.Println("Process 1 Both Locks acquired\n")
		lock1.Unlock()
		lock2.Unlock()
		fmt.Println("Process 1 Both Locks released\n")
	}
}	

func process2() {
	for {
		fmt.Println("Process 2 aquiring Lock 2\n")
		lock2.Lock()
		fmt.Println("Process 2 aquiring Lock 1\n")
		lock1.Lock()
		fmt.Println("Process 2 Both Locks acquired\n")
		lock1.Unlock()
		lock2.Unlock()
		fmt.Println("Process 2 Both Locks released\n")
	}
}	

func main() {
	go process1()
	go process2()
	time.Sleep(20 * time.Second)
	fmt.Println("Done")
}

