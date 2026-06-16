# Go Concurrency Assignment

This repository contains solutions for Go concurrency problems covering the core concepts of concurrent programming, synchronization, communication, cancellation, and one-time initialization.

## Problems

### Problem 1 - Concurrent Greetings

Print greetings concurrently using goroutines and synchronize execution using `sync.WaitGroup`.

### Problem 2 - Channels

Double numbers concurrently using goroutines and collect results through channels.

### Problem 3 - Concurrent Downloads

Simulate file downloads with random delays using goroutines and `sync.WaitGroup`.

### Problem 4 - Select and Timeout

Handle slow-running tasks using `select` and `time.After` to implement timeouts.

### Problem 5 - Mutex and Race Conditions

Protect shared data using `sync.Mutex` and ensure safe concurrent access.

### Problem 6 - Buffered Channel Logger

Build a simple logger using a buffered channel. Send multiple log messages into the channel and read them back in FIFO order without using goroutines.

### Problem 7 - Squares Generator with Channel Closing

Generate squares of numbers in a goroutine, send them through a channel, and use `close()` along with a `for range` loop to consume values safely.

### Problem 8 - Worker Pool / Order Processing System

Simulate processing 12 orders using 4 worker goroutines. Distribute work through a shared jobs channel and synchronize completion using `sync.WaitGroup`.

### Problem 9 - Ticker with Graceful Shutdown

Implement a ticker goroutine that periodically sends incrementing values through a channel and stops gracefully when a `done` channel is closed.

### Problem 10 - One-Time Configuration Loading

Use `sync.Once` to ensure configuration is loaded exactly once, even when multiple goroutines request it concurrently.

## Concepts Covered

### Concurrency Basics

* Goroutines
* `sync.WaitGroup`
* Channels
* Buffered Channels
* Channel Closing
* `for range` on Channels

### Coordination & Communication

* Worker Pools
* Producer-Consumer Pattern
* Shared Job Queues
* Done Channels
* Graceful Goroutine Shutdown

### Synchronization

* `sync.Mutex`
* Race Conditions
* `sync.Once`
* One-Time Initialization

### Control Flow

* `select`
* Timeouts
* Cancellation Signals

## Learning Outcomes

After completing these problems, you should understand:

* How to launch and synchronize goroutines.
* How channels enable safe communication between goroutines.
* The difference between buffered and unbuffered channels.
* How to distribute work among multiple workers.
* How to stop goroutines gracefully using cancellation signals.
* How to avoid race conditions using synchronization primitives.
* How to initialize shared resources exactly once with `sync.Once`.

## Author

Divakaran S
