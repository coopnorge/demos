package main

import "golang.org/x/sync/errgroup"

func processWithLimitErrGroup(items []string, maxWorkers int) {
	g := errgroup.Group{}
	g.SetLimit(maxWorkers)

	for _, item := range items {
		g.Go(func() error {
			process(item)
			return nil
		})
	}
	err := g.Wait()
	_ = err
}
