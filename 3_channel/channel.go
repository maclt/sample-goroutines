package main

import (
	"fmt"
	"sync"
)

func sendData(ch chan<- int, data int) {
	fmt.Println("Sending", data)
	ch <- data // Block until received
	fmt.Println("Finished sending", data)
}

func receiveData(ch <-chan int) {
	val := <-ch // Block until value is sent
	fmt.Println("Received", val)
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(2)
	go func() {
		defer wg.Done()
		sendData(ch, 10)
	}()
	go func() {
		defer wg.Done()
		receiveData(ch)
	}()

	wg.Wait()
	fmt.Println("Done")
}
