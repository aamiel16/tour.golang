package main

import (
	"fmt"
	"time"
	"math/rand"
)

type Order struct {
	orderId int
	order string
	processed bool
}

func producer(orders chan<- Order) {
	producerId := rand.Int()
	fmt.Printf("PROD [%d]: Producer spawned \n", producerId)

	for {
		// 1 Create a random order
		order := Order{
			orderId: rand.Int(),
			order: "Hamburger",
		}

		// 2 Add order to pending order channel
		fmt.Printf("PROD [%d]: Adding order (%d) \n", producerId, order.orderId)
		orders <- order

		// 3 Sleep
		time.Sleep(time.Second * time.Duration(rand.Int() % 5))
	}
}

func consumer(orders <-chan Order, processed chan<- Order) {
	consumerId := rand.Int()
	fmt.Printf("CONS [%d]: Consumer spawned \n", consumerId)

	// For any unprocessed order, process it
	for order := range orders {
		fmt.Printf("CONS [%d]: Processing order (%d) \n", consumerId, order.orderId)
		time.Sleep(time.Second * time.Duration(rand.Int() % 10))
		order.processed = true

		// Add order to processed channel
		processed <- order
	}
}

func main() {
	// Create pending orders, and processed orders channel
	orders := make(chan Order)
	processed := make(chan Order)

	// Spawn order producers
	go producer(orders)
	go producer(orders)

	// Spawn order consumers
	go consumer(orders, processed)
	go consumer(orders, processed)
	go consumer(orders, processed)

	for order := range processed {
		fmt.Printf("Order (%d) has been processed \n", order.orderId)
	}
}
