# Getting Started with stopwatch
Time counter, which can be used to time the execution of the code

```go
// new
sw := stopwatch.StartNew()

// sleep 1 second
time.Sleep(time.Second)

// print:finish，use：1000 ms
log.Println("finish，use：" + strconv.FormatInt(sw.ElapsedMilliseconds(), 10) + " ms")
```

```go
sw := stopwatch.StartNew()
// Reset Timer
sw.Restart()
// Continue timing
sw.Start
// Timer pause
sw.Stop
// Return to timed time
sw.ElapsedMilliseconds
```