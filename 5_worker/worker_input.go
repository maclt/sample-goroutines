package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func worker(id int, tasks <-chan string, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for t := range tasks {
		// Process the received string
		num, err := strconv.Atoi(t)
		if err != nil {
			results <- fmt.Sprintf("Worker %d: Invalid input '%s'", id, t)
			continue
		}
		results <- fmt.Sprintf("Worker %d: Task '%s' processed, result is %d", id, t, num*2)
	}
}

func main() {
	tasks := make(chan string, 10)
	results := make(chan string, 10)
	var wg sync.WaitGroup

	// Create 3 workers
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, tasks, results, &wg)
	}

	// Receive user input and send tasks
	go func() {
		defer close(tasks)
		fmt.Println("Enter numbers to process. Type 'exit' to quit.")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input := scanner.Text()
			if strings.ToLower(input) == "exit" {
				break
			}
			tasks <- input
		}
	}()

	// Wait for all workers to finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect and display results
	for r := range results {
		fmt.Println(r)
	}
}
