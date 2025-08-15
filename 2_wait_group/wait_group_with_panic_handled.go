package main

import (
	"fmt"
	"sync"
)

func task(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(name, "completed")
}

func taskWithPanic(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	panic("simulated panic")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	// this does not work!
	// https://moneyforward-dev.jp/entry/2020/09/09/goroutine-panic-handling/
	defer func() {
		if r3 := recover(); r3 != nil {
			fmt.Println("Recovered from panic:", r3)
		}
	}()

	go task("task 1", &wg)
	go task("task 2", &wg)
	go taskWithPanic("task 3", &wg)

	wg.Wait()

	fmt.Println("main function finished")
}
