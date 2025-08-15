package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 1; i < 4; i++ {
		fmt.Println(name, "running", i)
		time.Sleep(time.Second)
	}
}

func main() {
	task("task 1")
	task("task 2")
	fmt.Println("main function finished")
}
