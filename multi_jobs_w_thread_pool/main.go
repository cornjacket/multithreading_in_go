package main

// The purpose of this example is to give a worst case sequential processing of 10 jobs. The complementary example
// will use a worker pool to parallelize the job processing. This example should be somewhat deterministic in its
// processing times.

import (
	"fmt"
	"sync"
	"time"
)

var (
	waitGroup = sync.WaitGroup{}
)

const (
	maxTime = 4
	numThreads int = 4
)

func process_job(inputChannel chan int) {
	for job_id := range inputChannel {
        	time.Sleep(time.Duration(job_id%maxTime+1) * time.Second)
        	fmt.Printf("job: %d complete\n", job_id)
	}
	waitGroup.Done()
}

func main() {

	inputChannel := make(chan int, 1000)	// create buffered channel of int, 1000 deep
	for i:=0; i<numThreads; i++ {		// start up all the workers
		go process_job(inputChannel)
	}
	waitGroup.Add(numThreads)

	start := time.Now()
	for i:=0; i<10; i++ {
		inputChannel <- i
	}
	close(inputChannel)

	waitGroup.Wait()
	elapsed := time.Since(start)
        fmt.Printf("Main thread done. Total processing time: %s\n", elapsed)
}

