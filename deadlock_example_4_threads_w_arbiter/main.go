package main

// This program uses the arbiter pattern via single mutex and condition variable in order to allocate 2 shared resources (of 4) to
// 4 separate processes. Arbiter is implicitly implemeted inside each process. There is no arbiter process.

// process 1 wants to use resource 1 and 2
// process 2 wants to use resource 2 and 3
// process 3 wants to use resource 3 and 4
// process 4 wants to use resource 4 and 1
// last process runs every 100 ms and releases all resources 

import (
	"fmt"
	"sync"
	"time"
)

var (
	controller = sync.Mutex{}
	cond = sync.NewCond(&controller)
	
	resource1InUseBy = 0 // 0: no process, 1: process 1, 2: process 2, 3: process 3, 4: process 4 
	resource2InUseBy = 0 // 0: no process, 1: process 1, 2: process 2, 3: process 3, 4: process 4 
	resource3InUseBy = 0 // 0: no process, 1: process 1, 2: process 2, 3: process 3, 4: process 4 
	resource4InUseBy = 0 // 0: no process, 1: process 1, 2: process 2, 3: process 3, 4: process 4 
)

func resources1And2Busy() bool {
	return (resource1InUseBy!=0) || (resource2InUseBy!=0)
}

func resources2And3Busy() bool {
	return (resource2InUseBy!=0) || (resource3InUseBy!=0)
}

func resources3And4Busy() bool {
	return (resource3InUseBy!=0) || (resource4InUseBy!=0)
}

func resources4And1Busy() bool {
	return (resource4InUseBy!=0) || (resource1InUseBy!=0)
}


func process1() {
	for {
		controller.Lock()
		for resources1And2Busy() {
			cond.Wait()
		}
		fmt.Println("Process 1 aquiring Resource 1\n")
		resource1InUseBy = 1
		fmt.Println("Process 1 aquiring Resource 2\n")
		resource2InUseBy = 1
		fmt.Println("Process 1 Both Resources (1,2) acquired\n")
		controller.Unlock()
	}
}	

func process2() {
	for {
		controller.Lock()
		for resources2And3Busy() {
			cond.Wait()
		}
		fmt.Println("Process 2 aquiring Resource 2\n")
		resource2InUseBy = 2
		fmt.Println("Process 2 aquiring Resource 3\n")
		resource3InUseBy = 2
		fmt.Println("Process 2 Both Resources (2,3) acquired\n")
		controller.Unlock()
	}
}	

func process3() {
	for {
		controller.Lock()
		for resources3And4Busy() {
			cond.Wait()
		}
		fmt.Println("Process 3 aquiring Resource 3\n")
		resource3InUseBy = 3
		fmt.Println("Process 3 aquiring Resource 4\n")
		resource4InUseBy = 3
		fmt.Println("Process 3 Both Resources (3,4) acquired\n")
		controller.Unlock()
	}
}	

func process4() {
	for {
		controller.Lock()
		for resources4And1Busy() {
			cond.Wait()
		}
		fmt.Println("Process 4 aquiring Resource 4\n")
		resource4InUseBy = 4
		fmt.Println("Process 4 aquiring Resource 1\n")
		resource1InUseBy = 4
		fmt.Println("Process 4 Both Resources (4,1) acquired\n")
		controller.Unlock()
	}
}	

func releaseProcess() {
	for {
		controller.Lock()
		if resource1InUseBy != 0 {
			resource1InUseBy = 0
			fmt.Println("Resource 1 released\n")
		}
		if resource2InUseBy != 0 {
			resource2InUseBy = 0
			fmt.Println("Resource 2 released\n")
		}
		if resource3InUseBy != 0 {
			resource3InUseBy = 0
			fmt.Println("Resource 3 released\n")
		}
		if resource4InUseBy != 0 {
			resource4InUseBy = 0
			fmt.Println("Resource 4 released\n")
		}
		cond.Broadcast()
		controller.Unlock()
		time.Sleep(10 * time.Millisecond)
	}
}	
func main() {
	go process1()
	go process2()
	go process3()
	go process4()
	go releaseProcess()
	time.Sleep(20 * time.Second)
	fmt.Println("Done")
}

