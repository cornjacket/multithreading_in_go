package main

// This program uses the arbiter pattern via single mutex and condition variable in order to allocate 2 mutually exclusive resources to
// 2 separate processes. Arbiter is implicitly implemeted inside each process. There is no arbiter process.

import (
	"fmt"
	"sync"
	"time"
)

var (
	controller = sync.Mutex{}
	cond = sync.NewCond(&controller)
	
	resource1InUseBy = 0 // 0: no process, 1: process 1, 2: process 2 
	resource2InUseBy = 0 // 0: no process, 1: process 1, 2: process 2 
)

func resourcesNotFree() bool {
	return (resource1InUseBy!=0) || (resource2InUseBy!=0)
}


func process1() {
	for {
		controller.Lock()
		for resourcesNotFree() {
			cond.Wait()
		}
		fmt.Println("Process 1 aquiring Resource 1\n")
		resource1InUseBy = 1
		fmt.Println("Process 1 aquiring Resource 2\n")
		resource2InUseBy = 1
		fmt.Println("Process 1 Both Resources acquired\n")
		resource1InUseBy = 0
		resource2InUseBy = 0
		fmt.Println("Process 1 Both resources released\n")
		cond.Broadcast()
		controller.Unlock()
	}
}	

func process2() {
	for {
		controller.Lock()
		for resourcesNotFree() {
			cond.Wait()
		}
		fmt.Println("Process 2 aquiring Resource 2\n")
		resource2InUseBy = 2
		fmt.Println("Process 2 aquiring Resource 1\n")
		resource1InUseBy = 2
		fmt.Println("Process 2 Both Resources acquired\n")
		resource1InUseBy = 0
		resource2InUseBy = 0
		fmt.Println("Process 2 Both resources released\n")
		cond.Broadcast()
		controller.Unlock()
	}
}	

func main() {
	go process1()
	go process2()
	time.Sleep(20 * time.Second)
	fmt.Println("Done")
}

