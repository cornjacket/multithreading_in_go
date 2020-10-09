package main

import (
	"log"
	"sync"
	"time"
)

// The idea in the program is to test the readwrite lock mechanism by having the writer set only one element in an array to 1
// while the other elements are 0. The readers will continually scan the array making sure that only one element is set to 1
// while the others remain 0. Each reader will count the number of failures it encounters and print at the end of the experiment.

var (
	state		[100]int
	writer_done = false
	reader_done	[4]bool
	rwLock = sync.RWMutex{} 
)

// writer assumes that the first element is 1 prior the writer_thread is called.
func writer_thread() {
	println("writer_thread start")
	for i := 0; i < 1000; i++ {
		old_index := i % 100 // 100 is the size of the state array 
		next_index := (i+1) % 100 // 100 is the size of the state array
		rwLock.Lock()
		state[old_index] = 0
		time.Sleep(1 * time.Millisecond)
		state[next_index] = 1
		rwLock.Unlock()
		if !only_one_asserted() {
			log.Fatal("writer_thread failed sanity check")	
		}
		//time.Sleep(1 * time.Millisecond)
	}
	println("writer_thread stop")
	writer_done = true
}

func reader_thread(id int) {
	println("reader_thread start")
	fails := 0 
	for i := 0; i < 1000; i++ {
		if !only_one_asserted() {
			fails++
		} 
		time.Sleep(1 * time.Millisecond)
	}
	print("reader_thread ")
	print(id)
	print(" stop. fails = ")
	print(fails)
	println()
	reader_done[id] = true
}

func only_one_asserted() bool {
	count := 0
	rwLock.RLock()
	for j := 0; j < 100; j++ {
		count += state[j]
	}
	rwLock.RUnlock()
	if count != 1 {
		return false	
	} 
	return true
}

func main() {
	state[0] = 1
	go writer_thread()
	for i := 0; i < 4; i++ {
		go reader_thread(i)
	}
	allDone := false
	for ; !allDone; {

		// here is a boolean expression	
		readers_done := reader_done[0] && reader_done[1] && reader_done[2] && reader_done[3]
		allDone = writer_done && readers_done
		time.Sleep(1000 * time.Millisecond)
	}
	println("main done")
}
