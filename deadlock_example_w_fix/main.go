package main

// This program will not encounter a deadlock situation because there has been a heirarchy imposed on the locks.
// Both processes will attempt to lock Lock1 prior to attempting to lock Lock2. This way, each process will
// either lock both Locks always or neither locks and a deadlock will not occur.

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
		fmt.Println("Process 2 aquiring Lock 1\n")
		lock1.Lock()
		fmt.Println("Process 2 aquiring Lock 2\n")
		lock2.Lock()
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

