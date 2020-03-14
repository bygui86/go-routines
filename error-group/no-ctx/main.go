package main

import (
	"fmt"
	"math/rand"

	"golang.org/x/sync/errgroup"
)

func main() {
	var errGroup errgroup.Group

	for i := 0; i < 10; i++ {
		jobID := i
		errGroup.Go(func() error {
			return job(jobID)
		})
	}
	
	err := errGroup.Wait()
	if err != nil {
		fmt.Println("Encountered error:", err)
	}
	fmt.Println("Successfully finished.")
}

func job(jobID int) error {
	if rand.Intn(12) == jobID {
		return fmt.Errorf("job %v failed", jobID)
	}
	
	fmt.Printf("job %v done.\n", jobID)
	return nil
}
