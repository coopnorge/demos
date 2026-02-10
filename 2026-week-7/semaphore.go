package main

func processWithLimitSemaphore(items []string, maxWorkers int) {
	// Buffered capacity sets the maximum number of concurrent workers
	sem := make(chan struct{}, maxWorkers)

	for _, item := range items {
		sem <- struct{}{}
		go func(s string) {
			defer func() { <-sem }()
			process(s)
		}(item)
	}

	for i := 0; i < maxWorkers; i++ {
		sem <- struct{}{}
	}
}
