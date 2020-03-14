package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	errGroup, ctx := errgroup.WithContext(newCtx())

	for i := 0; i < 10; i++ {
		jobID := i
		errGroup.Go(func() error {
			return jobWithCtx(ctx, jobID)
		})
	}
	
	err := errGroup.Wait()
	if err != nil {
		fmt.Println("Encountered error:", err)
	}
	fmt.Println("Successfully finished.")
}

func newCtx() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	go waitingOsSignal(cancel)
	return ctx
}

func waitingOsSignal(cancel context.CancelFunc) {
	sCh := make(chan os.Signal, 1)
	signal.Notify(sCh, syscall.SIGINT, syscall.SIGTERM)
	<-sCh
	cancel()
}

func jobWithCtx(ctx context.Context, jobID int) error {
	select {
	case <-ctx.Done():
		fmt.Printf("context cancelled job %v terminting\n", jobID)
		return nil
	case <-time.After(time.Duration(rand.Intn(3)) * time.Second):
		// no-op
	}
	
	if rand.Intn(12) == jobID {
		fmt.Printf("Job %v failed.\n", jobID)
		return fmt.Errorf("job %v failed", jobID)
	}

	fmt.Printf("Job %v done.\n", jobID)
	return nil
}
