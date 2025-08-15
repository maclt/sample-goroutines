# Worker Input Example

This example demonstrates a worker pool pattern in Go, where user input is processed concurrently by multiple workers.

## Component Dependencies

- **main()**
  - Creates `tasks` and `results` channels.
  - Initializes a `sync.WaitGroup` for worker synchronization.
  - Launches 3 worker goroutines, each running the `worker` function.
  - Starts a goroutine to read user input and send tasks to the `tasks` channel.
  - Starts a goroutine to wait for all workers to finish and then close the `results` channel.
  - Collects and prints results from the `results` channel.

- **worker()**
  - Receives tasks from the `tasks` channel.
  - Processes each task (expects a string representing a number).
  - Sends the result or error message to the `results` channel.
  - Notifies the `WaitGroup` when done.

- **Input goroutine**
  - Reads user input from `os.Stdin` using `bufio.Scanner`.
  - Sends each input to the `tasks` channel until "exit" is entered.
  - Closes the `tasks` channel when done.

- **WaitGroup goroutine**
  - Waits for all workers to finish (`wg.Wait()`).
  - Closes the `results` channel.

### Dependency Flow

- Workers depend on the `tasks` channel for input and the `results` channel for output.
- The main function depends on the `results` channel for collecting output.
- The input goroutine depends on user input and sends data to the `tasks` channel.
- The WaitGroup ensures all workers finish before closing the `results` channel.

---

This structure allows for efficient, concurrent processing of user input using a pool of worker goroutines.
