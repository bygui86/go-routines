package main

import (
	"fmt"
	"sync"
)

var i int // i == 0

func main() {
	var wg sync.WaitGroup
	
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	
	// wait until all 1000 gorutines are done
	wg.Wait()
	
	// value of i should be 1000, but actually won't because of race condition between goroutines
	fmt.Println("value of i after 1000 operations is", i)
}

// goroutine increment global variable i
func worker(wg *sync.WaitGroup) {
	i = i + 1
	wg.Done()
}
