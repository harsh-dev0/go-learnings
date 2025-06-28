```markdown
## Understanding `defer` in Go

In Go, the `defer` keyword is used to schedule a function call to run after the surrounding function finishes, regardless of whether it exits normally or due to a panic.

For example:

```go
func sayHello(wg *sync.WaitGroup) {
    defer wg.Done()
    // function body
}
```

Here, `defer wg.Done()` ensures that `wg.Done()` will be called after `sayHello` completes. This is particularly useful for cleanup tasks or to guarantee that `wg.Done()` is always executed, even if the function returns early or encounters an error.

In the context of goroutines, using `defer wg.Done()` helps signal to the `WaitGroup` that the goroutine has finished its work.
```