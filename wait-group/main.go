package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("main() started")
	var wg sync.WaitGroup // create waitgroup (empty struct)
	
	for i := 1; i <= 3; i++ {
		wg.Add(1) // increment counter
		go service(&wg, i)
	}
	
	wg.Wait() // blocks here
	fmt.Println("main() stopped")
}

func service(wg *sync.WaitGroup, instance int) {
	defer wg.Done() // decrement counter
	time.Sleep(2 * time.Second)
	fmt.Println("Service called on instance", instance)
}
