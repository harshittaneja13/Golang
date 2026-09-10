package main

import (
	"fmt"
	"sync"
)

type Job struct {
	ID int
}

// func worker(id int, jobsChan chan Job, wg *sync.WaitGroup){
// 	defer wg.Done();
// 	for job := range jobsChan {
// 		fmt.Println("Worker ", id, "Processing job: ", job.ID)
// 	}
// }


func test(worker func(int, chan Job, *sync.WaitGroup)) {
	var wg sync.WaitGroup
	const numWorkers = 100
	jobChan := make(chan Job, 1000)
	
	// start the workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i+1, jobChan, &wg)
	}


	// Produce 1 million jobs
	for i := 1; i <= 1000000; i++ {
		jobChan <- Job{ID: i}
	}

	close(jobChan)

	wg.Wait()
	fmt.Println("All jobs processed")
}