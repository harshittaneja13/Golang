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
	mu sync.Mutex
} 

var (
	totalUpdates  int
	updateMutex sync.Mutex 
)

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

func processOrders(orders []*Order){
	for _, order := range orders {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		fmt.Println("Processing order: ", order.ID)
		// order.Status = "Processed"
	}
}

func updateOrderStatus(order *Order){
	order.mu.Lock()
	time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
	status := []string {
		"Delivered", "Shipped", "Processing",
	}[rand.Intn(3)]
	order.Status = status
	fmt.Println("Updated order: ", order.ID, "status: ", status)
	order.mu.Unlock() 

	updateMutex.Lock()
	defer updateMutex.Unlock()
	currentupdate := totalUpdates
	time.Sleep(5 * time.Millisecond) 
	totalUpdates = currentupdate+1
	
} 

func reportOrderStatus(orders []*Order){
	// for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("\n --------- Order Status Report: ---------")
		for _, order := range orders {
			fmt.Println("Order: ", order.ID, "status: ", order.Status)
		}
		fmt.Println("--------------------------------\n")
	// }
}

func main() {
	orders := generateOrders(20)
	// var wg sync.WaitGroup
	wg := sync.WaitGroup{}
	wg.Add(3)
	// go func(){
	// 	defer wg.Done()
	// 	processOrders(orders)
	// }()
	for i:=0;i<3;i++ {
		go func(){
			defer wg.Done()
			for _, order := range orders {  
				updateOrderStatus(order)
			}
		}()
	}
	

	wg.Wait()
	reportOrderStatus(orders)
	
	fmt.Println("All tasks completed successfully!") 
	fmt.Println("Total updates: ", totalUpdates)
}
