package main

import (
	"math/rand"
	"time"
)

// The idea in the program is to experiment with starting a pool of workers and waiting for them to complete.
// Since there is no synchronization between the main thread and the worker threads. The main thread must wait the
// worst case completion time for the experiment to be completed.
// The waitgroup version of this program will be faster on average since there is communication between the workers
// and the main thread when the workers complete.

// writer assumes that the first element is 1 prior the writer_thread is called.
func worker_thread(id int, max_time int) {
	print("worker_thread ")
	print(id)
	print(" begin.")
	println()
	time.Sleep(time.Duration(rand.Intn(max_time)) * time.Second)
	print("worker_thread ")
	print(id)
	print(" end.")
	println()
}

func main() {
	max_time := 10
	for i := 0; i < 10; i++ {
		go worker_thread(i, max_time)
	}
	time.Sleep(time.Duration(max_time) * time.Second)
	println("main done")
}
