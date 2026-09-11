package main

import (
	"fmt"
	"sync"
	"time"
)

// func worker(url string, wg *sync.WaitGroup, resultChan chan string){
// 	defer wg.Done()
// 	time.Sleep(50*time.Millisecond)
// 	fmt.Printf("Image Processed: %s \n", url)
// 	resultChan<-url
// }

func worker(jobsChan chan string, wg *sync.WaitGroup, resultChan chan string){
	defer wg.Done()

	time.Sleep(50*time.Millisecond)
	
	for job := range jobsChan{
		// fmt.Printf("Job Completed: %s \n", job)
		resultChan <- job
	}

	fmt.Println("Worker Finished")
	
}


 
func main(){
	jobs := []string{
		"Image1.png",
		"Image2.png",
		"Image3.png",
		"Image4.png",
		"Image5.png",
		"Image6.png",
		"Image7.png",
		"Image8.png",
		"Image9.png",
		"Image10.png",
		"Image11.png",
		"Image12.png",
		"Image13.png",
		"Image14.png",
		"Image15.png",
		"Image16.png",
		"Image17.png",
		"Image18.png",
		"Image19.png",
		"Image20.png",
	}
	
	startTime := time.Now()
	var wg sync.WaitGroup

	totalWorkers := 10;
	
	jobsChan := make(chan string, 20)
	resultChan := make(chan string, 20) 

	
	for i:=0;i<totalWorkers;i++ {
		wg.Add(1)
		go worker(jobsChan, &wg, resultChan)
	}
	
	// producer
	for _, job := range jobs {
		jobsChan <- job
	}
	close(jobsChan)


	// wg.Add(5)
	// go worker("Image1.png", &wg, resultChan)
	// go worker("Image2.png", &wg, resultChan)
	// go worker("Image3.png", &wg, resultChan)
	// go worker("Image4.png", &wg, resultChan)
	// go worker("Image5.png", &wg, resultChan)

	

	go func(){
		wg.Wait()
		close(resultChan)
	}()
	
	for val := range resultChan {
		fmt.Printf("Received : %s \n", val)
	}
	// defer fmt.Println(1)
	// defer fmt.Println(2)
	// defer fmt.Println(3)
	
	fmt.Printf("It took %s ms. \n", time.Since(startTime))

}