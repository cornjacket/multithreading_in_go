package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	waitgroup = sync.WaitGroup{}
)

// The idea in the program is to experiment with starting a pool of workers and waiting for them to complete.
// The waitgroup version of this program will be faster on average since there is communication between the workers
// using a waitgroup. The main thread terminates immediately after the last worker completes where as in the
// non-wait group version the main thread waits the max time that any of the threads will run.

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
	waitgroup.Done()
}

func main() {
	start := time.Now()
	max_time := 10
	for i := 0; i < 10; i++ {
		waitgroup.Add(1)
		go worker_thread(i, max_time)
	}
	waitgroup.Wait()
	println("main done")
	elapsed := time.Since(start)
	fmt.Printf("Processing took %s\n", elapsed)
}
