package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

// expensiveFunction simulates an expensive call, like a database connection, or HTTP request.
func expensiveFunction(ctx context.Context, i int, duration time.Duration) error {
	ch := time.After(duration)
	select {
	case <-ch:
		log.Printf("Done running expensive function %d", i)
		return nil
	case <-ctx.Done():
		log.Printf("Context was cancelled when running expensive function %d", i)
		return ctx.Err()
	}
}

// expensiveFailingFunction simulates an expensive call, like a database connection, or HTTP request, that will always fail.
func expensiveFailingFunction(ctx context.Context, i int, duration time.Duration) error {
	ch := time.After(duration)
	select {
	case <-ch:
		log.Printf("Done running expensive function %d, failing intentionally", i)
		return fmt.Errorf("Intentionally failed expensive function %d", i)
	case <-ctx.Done():
		log.Printf("Context was cancelled when running expensive function %d", i)
		return ctx.Err()
	}
}
