package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	mu    sync.Mutex
	count int
}

func (sc *SafeCounter) increment() {
	sc.mu.Lock()
	sc.count++
	sc.mu.Unlock()
}

func main() {
	sc := &SafeCounter{}
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			sc.increment()
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			sc.increment()
		}
	}()

	wg.Wait()
	fmt.Println("Final count:", sc.count)
}
