package main

import (
	"fmt"
	"sync"
)

const (
	Reset  = "\x1b[0m" // Reset color
	Red    = "\x1b[31m"
	Green  = "\x1b[32m"
	Yellow = "\x1b[33m"
	Blue   = "\x1b[34m"
)

type UnsafeCounter struct {
	// mu    sync.Mutex
	count int
}

func (sc *UnsafeCounter) increment() {
	// sc.mu.Lock()
	sc.count++
	// sc.mu.Unlock()
}

func trial() {
	sc := &UnsafeCounter{}
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
	if sc.count != 2000 {
		fmt.Println("Final count:", Yellow+fmt.Sprint(sc.count)+Reset)
	} else {
		fmt.Println("Final count:", sc.count)
	}
}

func main() {
	for i := 0; i < 20; i++ {
		trial()
	}

	fmt.Println("Did you find any count less than 2000?")
}
