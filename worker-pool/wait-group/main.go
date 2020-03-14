package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Println("[main] started")
	
	var wg sync.WaitGroup
	
	tasks := make(chan int, 10)
	results := make(chan int, 10)
	
	fmt.Println("[main] launch worker goroutines")
	for i := 0; i < 3; i++ {
		go sqrWorker(&wg, tasks, results, i)
		wg.Add(1)
	}
	
	fmt.Println("[main] pass tasks to tasks-channel")
	for i := 0; i < 5; i++ {
		tasks <- i * 2 // non-blocking as buffer capacity is 10
	}
	
	fmt.Printf("[main] %d running goroutines\n", runtime.NumGoroutine())
	
	fmt.Println("[main] close tasks-channel")
	close(tasks)
	
	fmt.Println("[main] wait until all workers do their job")
	wg.Wait()
	
	fmt.Println("[main] receive results on result-channel from all workers")
	for i := 0; i < 5; i++ {
		result := <-results // blocking because buffer is empty
		fmt.Println("[main] Result", i, ":", result)
	}
	
	fmt.Println("[main] stopped")
}

// worker than make squares
func sqrWorker(wg *sync.WaitGroup, tasks <-chan int, results chan<- int, instance int) {
	fmt.Printf("[worker-%d] started\n", instance)
	
	for num := range tasks {
		time.Sleep(time.Millisecond) // simulating blocking task
		fmt.Printf("[worker %d] send result back\n", instance)
		results <- num * num
	}
	
	// once tasks-channel will be closed
	fmt.Printf("[worker-%d] stopped\n", instance)
	
	// done with worker
	wg.Done()
}
