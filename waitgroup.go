package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(jobID int) {
			defer wg.Done()
			fmt.Printf("This is job: %v\n", jobID)
		}(i)
	}
	wg.Wait()
}
