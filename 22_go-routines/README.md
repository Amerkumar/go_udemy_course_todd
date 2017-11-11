#Concurrency vs Parallelism
- Concurrency is the composition of independently executing processes, while parallelism is the simultaneous execution of (possibly related) computations. Concurrency is about dealing with lots of things at once. Parallelism is about doing lots of things at once.
- See the slides. 
- Function init that run before main for a single time.
- Race conditions when you have things concurrently and different processors are trying to access the same variable and you can have some overriding.
```go
     go run -race main.go
```
- Mutex stands for **Mutual Exclusion Object.** is a program object that is created so that multiple program thread can take turns sharing the same resource such as access to file.
- Atomicity is another way to prevent race conditions. You have to indentify very small part of code that should be accessed by only one thread.


#Channels
- Do not communicate by sharing memory, share memory by communicating. 
- Concurrency is not parallelism.
- Use channels rather than mutex.
- Channel are way to communicate between goroutines.
