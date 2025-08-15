package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 1; i < 4; i++ {
		fmt.Println(name, "running", i)
		time.Sleep(time.Microsecond * 500)
	}
}

func main() {
	go task("task 1")
	go task("task 2")
	fmt.Println("main function finished")
}
