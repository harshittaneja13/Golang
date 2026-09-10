package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Order struct {
	ID int
	Status string
	// mu sync.Mutex
} 


func generateOrders(count int) []*Order {
	orders := make([]*Order, count)
	for i:=0;i<count;i++ {
		orders[i] = &Order{
			ID: i+1,
			Status: "Pending",
		}
	}
	return orders

}

func processOrders(inChan <-chan *Order, outChan chan<- *Order, wg *sync.WaitGroup){  // receive only channel and send only channel
	defer func(){
		wg.Done()
		close(outChan)
	}()  

	// it loops over orderCHan unitil it is closed 
	for order := range inChan {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		order.Status = "Processed"
		outChan <- order
	}
}

func main() {
	
	// var wg sync.WaitGroup
	wg := sync.WaitGroup{}
	wg.Add(3)

	orderChan := make(chan *Order, 20)
	processedChan := make(chan *Order, 20)

	go func(){
		defer wg.Done()
		for _, order := range generateOrders(20){
			orderChan <- order
		}
		close(orderChan)
		fmt.Println("Done with generating orders. ", )
	}()

	go processOrders(orderChan, processedChan, &wg)
	
	go func ()  {
		defer wg.Done()

		for {
			select {
				case  processedOrder, ok := <-processedChan: 
					if !ok {
						fmt.Println("Processing channel closed")
						return
					}
					fmt.Printf("Processed order %d with status: %s \n", processedOrder.ID, processedOrder.Status)
				case <-time.After(460*time.Millisecond):
					fmt.Println("Timeout for operations. ")
					return
			}
		}


	}()

	

	wg.Wait()
	
	fmt.Println("All tasks completed successfully!") 
}
