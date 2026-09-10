package main

import (
	"fmt"
	"sync"
	"time"
)

type Result struct {
	Value string
	Err   error
}

func worker(url string, wg *sync.WaitGroup, resultChan chan Result) {
	defer wg.Done()
	time.Sleep(50 * time.Millisecond)
	fmt.Println("image processed: ", url)

	resultChan <- Result{Value: url, Err: nil}
}

func workerFromChannel(jobsChan chan string, wg *sync.WaitGroup, resultChan chan Result) {
	defer wg.Done()

	for job := range jobsChan {
		time.Sleep(50 * time.Millisecond)
		fmt.Println("image processed: ", job)
		resultChan <- Result{Value: job, Err: nil}
	}

	fmt.Print("Worker shutting down")
}

func main() {
	var wg sync.WaitGroup
	startTime := time.Now()
	// to return from goroutine we use channel
	resultChan := make(chan Result, 50)
	// Fan-out pattern: multiple goroutines processing images concurrently
	// wg.Add(7)
	// go worker("image1.jpg", &wg, resultChan)
	// go worker("image2.jpg", &wg, resultChan)
	// go worker("image3.jpg", &wg, resultChan)
	// go worker("image4.jpg", &wg, resultChan)
	// go worker("image5.jpg", &wg, resultChan)
	// go worker("image6.jpg", &wg, resultChan)
	// go worker("image7.jpg", &wg, resultChan)

	// wg.Wait()
	// close(resultChan)

	// // Fan-in pattern: collecting results from multiple goroutines
	// for res := range resultChan {
	// 	fmt.Println("result received: ", res.Value, "error: ", res.Err)
	// 	// error handling can be done here if res.Err is not nil
	// 	if res.Err != nil {
	// 		// handle error
	// 		// reprocess - queue - dead letter queue
	// 	}
	// }

	// fmt.Println("All images processed in: ", time.Since(startTime)) // processed concurrently

	jobs := []string{"image1.jpg", "image2.jpg", "image3.jpg", "image4.jpg", "image5.jpg", "image6.jpg", "image7.jpg", "image8.jpg", "image9.jpg", "image10.jpg",
		"image11.jpg", "image12.jpg", "image13.jpg", "image14.jpg", "image15.jpg", "image16.jpg", "image17.jpg", "image18.jpg", "image19.jpg", "image20.jpg"}

	// for _, job := range jobs {
	// 	wg.Add(1)
	// 	go worker(job, &wg, resultChan)     // we have to limit the number of concurrent goroutines to avoid overwhelming the system, if there are 10000 images to process then 10000 goroutines will run concurrently and it will overwhelm the system, so we can limit the number of concurrent goroutines by using a buffered channel or a semaphore pattern
	// }

	// wg.Wait()
	// close(resultChan)

	// for res := range resultChan {
	// 	fmt.Println("result received: ", res.Value, "error: ", res.Err)
	// 	if res.Err != nil {
	// 		// handle error
	// 		// reprocess - queue - dead letter queue
	// 	}
	// }

	// fmt.Println("All images processed in: ", time.Since(startTime)) // processed concurrently

	totalWorkers := 5                        // limit the number of concurrent goroutines to 5
	jobsChan := make(chan string, len(jobs)) // buffered channel to hold jobs

	// Start worker goroutines
	for i := 1; i <= totalWorkers; i++ {
		wg.Add(1)
		go workerFromChannel(jobsChan, &wg, resultChan)
		fmt.Println("Worker ", i, " started")
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// send the jobs to the jobs channel
	for i := 0; i < len(jobs); i++ {
		jobsChan <- jobs[i]
	}

	close(jobsChan)

	for result := range resultChan {
		fmt.Println("Job Completed: ", result.Value, "error: ", result.Err)
		if result.Err != nil {
			// handle error
			// reprocess - queue - dead letter queue
		}
	}

	fmt.Println("All images processed in: ", time.Since(startTime)) // processed concurrently
}
