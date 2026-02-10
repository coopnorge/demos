package main

func processWithLimitWorkerPool(items []string, maxWorkers int) {
	jobs := make(chan string, maxWorkers)
	doneCh := make(chan struct{})

	// Start workers
	for range maxWorkers {
		go func() {
			for {
				item, ok := <-jobs
				if !ok {
					// signal worker is done
					doneCh <- struct{}{}
					return
				}
				process(item)
			}
		}()
	}

	// Fill the channel with items
	for _, item := range items {
		jobs <- item
	}
	close(jobs) // Signal that no more items will be sent

	for range maxWorkers {
		// Wait for all workers to be done
		<-doneCh
	}
}
