# Demonstrates the simplicity of golang.org/x/sync/errgroup

This demo demonstrates a few different ways to run concurrent code in Go,
and in the end some of the additional features that errgroup brings to the table.

How to run the examples:

```bash
go run . exampleSynchronousCalls
go run . exampleNakedGoRoutines
go run . exampleNakedGoRoutinesWithChannels
go run . exampleNakedGoRoutinesWithWaitgroup
go run . exampleNakedGoRoutinesWithWaitgroupSingleError
go run . exampleNakedGoRoutinesWithWaitgroupSingleErrorCancelContext
go run . exampleErrgroup
go run . exampleErrgroupWithContext
go run . exampleErrgroupWithSemaphore
```
