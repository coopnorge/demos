package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

/*
go run main.go exampleSynchronousCalls
go run main.go exampleNakedGoRoutines
go run main.go exampleNakedGoRoutinesWithChannels
go run main.go exampleNakedGoRoutinesWithWaitgroup
go run main.go exampleNakedGoRoutinesWithWaitgroupSingleError
go run main.go exampleNakedGoRoutinesWithWaitgroupSingleErrorCancelContext
go run main.go exampleErrgroup
go run main.go exampleErrgroupWithContext
go run main.go exampleErrgroupWithSemaphore
*/

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	example := "exampleSynchronousCalls"
	if len(os.Args) >= 2 {
		example = os.Args[1]
	}

	switch example {
	case "exampleSynchronousCalls":
		exampleSynchronousCalls(ctx)
	case "exampleNakedGoRoutines":
		exampleNakedGoRoutines(ctx)
	case "exampleNakedGoRoutinesWithChannels":
		exampleNakedGoRoutinesWithChannels(ctx)
	case "exampleNakedGoRoutinesWithWaitgroup":
		exampleNakedGoRoutinesWithWaitgroup(ctx)
	case "exampleNakedGoRoutinesWithWaitgroupSingleError":
		exampleNakedGoRoutinesWithWaitgroupSingleError(ctx)
	case "exampleNakedGoRoutinesWithWaitgroupSingleErrorCancelContext":
		exampleNakedGoRoutinesWithWaitgroupSingleErrorCancelContext(ctx)
	case "exampleErrgroup":
		exampleErrgroup(ctx)
	case "exampleErrgroupWithContext":
		exampleErrgroupWithContext(ctx)
	case "exampleErrgroupWithSemaphore":
		exampleErrgroupWithSemaphore(ctx)
	default:
		log.Println("Unknown example")
		os.Exit(1)
	}

	log.Printf("Done running full example")
}

func exampleSynchronousCalls(ctx context.Context) {
	if err := expensiveFunction(ctx, 1, 1*time.Second); err != nil {
		log.Fatalf("Error running expensive call 1")
	}

	if err := expensiveFunction(ctx, 2, 1*time.Second); err != nil {
		log.Fatalf("Error running expensive call 2")
	}
}

// exampleNakedGoRoutines demonstrates how we can use naked go-routines to start multiple tasks in parallel
// NOTE: This example does NOT introduce a synchronization-point, and therefore has big problems.
func exampleNakedGoRoutines(ctx context.Context) {
	var err1, err2 error
	go func() {
		err1 = expensiveFunction(ctx, 1, 1*time.Second)
	}()
	go func() {
		err2 = expensiveFunction(ctx, 2, 1*time.Second)
	}()

	if err1 != nil {
		log.Fatalf("Error running expensive call 1")
	}
	if err2 != nil {
		log.Fatalf("Error running expensive call 2")
	}
}

// exampleNakedGoRoutinesWithChannels demonstrates how we can use naked go-routines to start multiple tasks in parallel
// NOTE: This example uses a channel to synchronize.
func exampleNakedGoRoutinesWithChannels(ctx context.Context) {
	ch := make(chan struct{})
	var err1, err2 error

	go func() {
		err1 = expensiveFunction(ctx, 1, 1*time.Second)
		ch <- struct{}{}
	}()
	go func() {
		err2 = expensiveFunction(ctx, 2, 1*time.Second)
		ch <- struct{}{}
	}()

	// Read twice from the channel. Will block until an empty struct is sent on the channel
	<-ch
	<-ch

	if err1 != nil {
		log.Fatalf("Error running expensive call 1")
	}
	if err2 != nil {
		log.Fatalf("Error running expensive call 2")
	}
}

// exampleNakedGoRoutinesWithWaitgroup demonstrates how we can use naked go-routines to start multiple tasks in parallel
// NOTE: This example uses a Waitgroup to synchronize.
func exampleNakedGoRoutinesWithWaitgroup(ctx context.Context) {
	wg := sync.WaitGroup{}
	var err1, err2 error

	// Add 2 to the Waitgroup, since we are going to start 2 goroutines
	wg.Add(2)

	go func() {
		defer wg.Done()
		err1 = expensiveFunction(ctx, 1, 1*time.Second)
	}()
	go func() {
		defer wg.Done()
		err2 = expensiveFunction(ctx, 2, 1*time.Second)
	}()

	// Wait until the Waitgroup's counter has been reduced to 0.
	wg.Wait()

	if err1 != nil {
		log.Fatalf("Error running expensive call 1")
	}
	if err2 != nil {
		log.Fatalf("Error running expensive call 2")
	}
}

// exampleNakedGoRoutinesWithWaitgroupSingleError demonstrates how we can use naked go-routines to start multiple tasks in parallel
// NOTE: This example uses a Waitgroup to synchronize, and only keeps 1 error
func exampleNakedGoRoutinesWithWaitgroupSingleError(ctx context.Context) {
	wg := sync.WaitGroup{}
	var firstErr error
	errOnce := sync.Once{}

	// Add 2 to the Waitgroup, since we are going to start 2 goroutines
	wg.Add(2)

	go func() {
		defer wg.Done()
		err := expensiveFunction(ctx, 1, 1*time.Second)
		if err != nil {
			errOnce.Do(func() {
				firstErr = err
			})
		}
	}()
	go func() {
		defer wg.Done()
		err := expensiveFunction(ctx, 2, 1*time.Second)
		if err != nil {
			errOnce.Do(func() {
				firstErr = err
			})
		}
	}()

	// Wait until the Waitgroup's counter has been reduced to 0.
	wg.Wait()

	if firstErr != nil {
		log.Fatalf("Error running an expensive call. We don't know explicitly which one failed: %v", firstErr.Error())
	}
}

// exampleNakedGoRoutinesWithWaitgroupSingleErrorCancelContext demonstrates how we can use naked go-routines to start multiple tasks in parallel
// NOTE: This example uses a Waitgroup to synchronize, and only keeps 1 error, and also cancels the context on the first err
func exampleNakedGoRoutinesWithWaitgroupSingleErrorCancelContext(ctx context.Context) {
	gctx, cancel := context.WithCancel(ctx)
	defer cancel()

	wg := sync.WaitGroup{}
	var firstErr error
	errOnce := sync.Once{}

	// Add 2 to the Waitgroup, since we are going to start 2 goroutines
	wg.Add(2)

	go func() {
		defer wg.Done()
		err := expensiveFailingFunction(gctx, 1, 1*time.Second) // Fail after 1 second
		if err != nil {
			errOnce.Do(func() {
				firstErr = err
				cancel()
			})
		}
	}()
	go func() {
		defer wg.Done()
		err := expensiveFunction(gctx, 2, 20*time.Second) // Succeed after 20 seconds
		if err != nil {
			errOnce.Do(func() {
				firstErr = err
				cancel()
			})
		}
	}()

	// Wait until the Waitgroup's counter has been reduced to 0.
	wg.Wait()

	if firstErr != nil {
		log.Fatalf("Error running an expensive call. We don't know explicitly which one failed: %v", firstErr.Error())
	}
}

// exampleErrgroup demonstrates how we can use naked go-routines to start multiple tasks in parallel
func exampleErrgroup(ctx context.Context) {
	g := errgroup.Group{}

	g.Go(func() error {
		return expensiveFunction(ctx, 1, 1*time.Second)
	})

	g.Go(func() error {
		return expensiveFunction(ctx, 2, 1*time.Second)
	})

	// Wait until the Errgroup's counter has been reduced to 0.
	err := g.Wait()
	if err != nil {
		log.Fatalf("Error running an expensive call. We don't know explicitly which one failed: %v", err.Error())
	}
}

// exampleErrgroupWithContext demonstrates how to cancel the errgroup-context once a function fails.
func exampleErrgroupWithContext(ctx context.Context) {
	// By using "WithContext", the first error will also cancel the gctx
	g, gctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return expensiveFailingFunction(gctx, 1, 1*time.Second) // Fail after 1 second
	})

	g.Go(func() error {
		return expensiveFunction(gctx, 2, 20*time.Second) // Succeed after 20 seconds
	})

	// Wait until the Errgroup's counter has been reduced to 0.
	err := g.Wait()
	if err != nil {
		log.Fatalf("Error running an expensive call. We don't know explicitly which one failed: %v", err.Error())
	}
}

// exampleErrgroupWithSemaphore demonstrates how to cancel the errgroup-context once a function fails.
func exampleErrgroupWithSemaphore(ctx context.Context) {
	// By using "WithContext", the first error will also cancel the gctx
	g, gctx := errgroup.WithContext(ctx)

	// Set a limit of 2 active goroutines (semaphore), to prevent being rate limited by external service
	g.SetLimit(2)

	for i := 0; i < 10; i++ {
		g.Go(func() error {
			return expensiveFunction(gctx, i, 1*time.Second) // Succeed after 1 second
		})
	}

	// Wait until the Errgroup's counter has been reduced to 0.
	err := g.Wait()
	if err != nil {
		log.Fatalf("Error running an expensive call. We don't know explicitly which one failed: %v", err.Error())
	}
}
