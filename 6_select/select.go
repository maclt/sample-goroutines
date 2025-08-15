package main

import (
	"fmt"
	"time"

// task1 sends a message to the channel every second
func task1(ch chan string) {
	for {
		time.Sleep(1 * time.Second)
		ch <- "Task 1 completed"
	}
}

// task2 sends a message to the channel every two seconds
func task2(ch chan string) {
	for {
		time.Sleep(2 * time.Second)
		ch <- "Task 2 completed"
	}
}

func main() {
	// Create 2 channels
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Start 2 goroutines
	go task1(ch1)
	go task2(ch2)

	// Wait for messages from both channels
	for i := 0; i < 5; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}
}
