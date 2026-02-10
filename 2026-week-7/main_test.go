package main

import (
	"fmt"
	"testing"
)

const (
	smallSize = 5_000
	bigSize   = 2_000_000
	workers   = 2
)

func BenchmarkSemaphore_Small(b *testing.B) {
	items := make([]string, 0, smallSize)
	for i := range smallSize {
		items = append(items, fmt.Sprintf("item%d", i))
	}
	for b.Loop() {
		processWithLimitSemaphore(items, workers)
	}
}

func BenchmarkErrGroup_Small(b *testing.B) {
	items := make([]string, 0, smallSize)
	for i := range smallSize {
		items = append(items, fmt.Sprintf("item%d", i))
	}
	for b.Loop() {
		processWithLimitErrGroup(items, workers)
	}
}

func BenchmarkWorkerPool_Small(b *testing.B) {
	items := make([]string, 0, smallSize)
	for i := range smallSize {
		items = append(items, fmt.Sprintf("item%d", i))
	}
	for b.Loop() {
		processWithLimitWorkerPool(items, workers)
	}
}

func BenchmarkSemaphore_Big(b *testing.B) {
	items := make([]string, 0, bigSize)
	for i := range bigSize {
		items = append(items, fmt.Sprintf("item%d", i))
	}
	for b.Loop() {
		processWithLimitSemaphore(items, workers)
	}
}

func BenchmarkErrGroup_Big(b *testing.B) {
	items := make([]string, 0, bigSize)
	for i := range bigSize {
		items = append(items, fmt.Sprintf("item%d", i))
	}
	for b.Loop() {
		processWithLimitErrGroup(items, workers)
	}
}

func BenchmarkWorkerPool_Big(b *testing.B) {
	items := make([]string, 0, bigSize)
	for i := range bigSize {
		items = append(items, fmt.Sprintf("item%d", i))
	}
	for b.Loop() {
		processWithLimitWorkerPool(items, workers)
	}
}
