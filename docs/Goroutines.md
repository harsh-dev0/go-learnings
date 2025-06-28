# Goroutines in Go

Go provides a lightweight way to achieve concurrency using **goroutines**.

## What is Concurrency?

**Concurrency** is the ability of a program to deal with many tasks at once. It doesn't mean tasks are running simultaneously (that's **parallelism**), but rather that they are making progress independently.

> **Concurrency ≠ Parallelism**  
> - *Concurrency* is about structuring a program to handle multiple tasks at once.  
> - *Parallelism* is about running multiple tasks at the same time, typically on multiple CPUs.

## Starting a Goroutine

A goroutine is a function running independently in the same address space. You start a goroutine by prefixing a function call with the `go` keyword:


**Usage**:

```go
go functionName()
```

**Note**: Runs concurrently. Main function may exit before goroutines finish unless synchronized.

---

## WaitGroup

**Purpose**: Wait for multiple goroutines to finish.

### Key Methods:

* `wg.Add(n)` – set number of goroutines to wait for
* `wg.Done()` – called by each goroutine when finished
* `wg.Wait()` – blocks until all goroutines call `Done()`

**Usage**:

```go
var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    // work
}()
wg.Wait()
```

---

## Mutex

**Purpose**: Prevent race conditions when writing to shared data.

### Key Methods:

* `mu.Lock()` – acquire lock (exclusive)
* `mu.Unlock()` – release lock

**Usage**:

```go
var mu sync.Mutex

mu.Lock()
// modify shared data
mu.Unlock()
```

---

## RWMutex

**Purpose**: Optimize concurrent access – allow multiple readers, only one writer.

### Key Methods:

* `mu.RLock()` – shared read lock
* `mu.RUnlock()` – release read lock
* `mu.Lock()` – exclusive write lock
* `mu.Unlock()` – release write lock

**Usage**:

```go
var mu sync.RWMutex

mu.RLock()
// read shared data
mu.RUnlock()

mu.Lock()
// write shared data
mu.Unlock()
```

---

## Gotchas and Tips

* Always call `Unlock()` or `RUnlock()` after locking.
* Use `defer` immediately after lock to avoid forgetting.
* `Mutex` blocks all access. `RWMutex` allows concurrent reads.
* Use `RWMutex` when reads are frequent, writes are rare.
* Avoid data races using `go run -race`.
* Never mix plain reads/writes with concurrent goroutines without sync.

---

## Why Not Use Lock for Everything? When to Use RLock?

Using a regular `Lock` for every access (read or write) to shared data can severely limit concurrency, because only one goroutine can access the data at a time—even if many goroutines only need to read. This can become a bottleneck in read-heavy workloads.

### When to Use RLock/RUnlock

- Use `RLock`/`RUnlock` (from `sync.RWMutex`) when you have multiple goroutines that only need to read shared data.
- Multiple readers can hold the lock simultaneously, improving performance.
- If any goroutine needs to write, it must use `Lock`/`Unlock`, which will block all readers and writers until the write is done.

### Example Scenario

Suppose you have a shared map that is read frequently but only occasionally written to. Using `RWMutex` with `RLock` for reads and `Lock` for writes allows many readers to access the map at once, but ensures exclusive access for writers.

### Summary

- Use `Lock`/`Unlock` for exclusive (write) access.
- Use `RLock`/`RUnlock` for concurrent read-only access.
- Prefer `RWMutex` when reads greatly outnumber writes to maximize concurrency and performance.
- Avoid using only `Lock` for everything, as it unnecessarily serializes all access.

# Channels in Goroutines

## What are Channels?

**Channels** are Go's way to communicate safely between goroutines. They provide a typed conduit for sending and receiving values, ensuring synchronization.

### Declaring and Creating Channels

```go
ch := make(chan int) // unbuffered channel of int
```

### Sending and Receiving

```go
ch <- 42      // send value to channel
value := <-ch // receive value from channel
```

- Sending blocks until another goroutine is ready to receive (and vice versa).

### Example: Basic Channel Usage

```go
ch := make(chan string)

go func() {
    ch <- "hello"
}()

msg := <-ch
fmt.Println(msg) // Output: hello
```

### Buffered Channels

Channels can have a buffer, allowing sends to proceed without waiting (until the buffer is full):

```go
ch := make(chan int, 2)
ch <- 1
ch <- 2
// ch <- 3 // would block if buffer is full
```

### Closing Channels

Close a channel to signal no more values will be sent:

```go
close(ch)
```

Receiving from a closed channel yields zero value after all data is consumed.

### Range Over Channels

You can use `for range` to receive values until a channel is closed:

```go
for v := range ch {
    fmt.Println(v)
}
```

### Select Statement

`select` lets a goroutine wait on multiple channel operations:

```go
select {
case v := <-ch1:
    fmt.Println("Received", v)
case ch2 <- 10:
    fmt.Println("Sent 10")
default:
    fmt.Println("No communication")
}
```

### Summary

- Use channels to safely share data between goroutines.
- Unbuffered channels synchronize sender and receiver.
- Buffered channels allow limited asynchronous communication.
- Always close channels when done sending.
- Use `select` for handling multiple channels.