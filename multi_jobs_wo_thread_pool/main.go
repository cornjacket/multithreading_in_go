package main

import (
	"fmt"
	"time"
)

const max_time = 4

func process_job(job_id int) {
        time.Sleep(time.Duration(job_id%max_time+1) * time.Second)
        fmt.Printf("job: %d complete\n", job_id)
}

func main() {
	start := time.Now()
	for i:=0; i<10; i++ {
		process_job(i)
	}
	elapsed := time.Since(start)
        fmt.Printf("Main thread done. Total processing time: %s\n", elapsed)
}

