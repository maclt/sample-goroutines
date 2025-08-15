# Goroutines Tutorial

This repository is a step-by-step tutorial for learning and experimenting with Go's concurrency primitives, especially goroutines. Each folder and file demonstrates a specific concept, building up your understanding of concurrent programming in Go.

## Table of Contents

1. **0_serialize/** — Basic synchronization primitives in Go.
2. **1_goroutine/** — Introduction to goroutines: how to start and use them.
3. **2_wait_group/** — Using `sync.WaitGroup` to wait for goroutines to finish, including panic handling.
4. **3_channel/** — Communicating between goroutines using channels.
5. **4_mutex/** — Protecting shared data with mutexes, and what happens without them.
6. **5_worker/** — Building worker pools and managing concurrent tasks.
7. **6_select/** — Using the `select` statement to handle multiple channel operations.

## How to Use

1. Clone this repository.
2. Explore each folder in order. Each step builds on the previous one.
3. Run the Go files to see the output and experiment with the code.

```sh
go run <filename.go>
```

4. Ask LLM chat to explain each sample code.

## License

MIT