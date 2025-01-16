package main

import (
	"runtime"
	"sort"
	"sync"
)

type ParallelPrimeChecker struct {
	numWorkers int
}

func NewParallelPrimeChecker() *ParallelPrimeChecker {
	return &ParallelPrimeChecker{
		numWorkers: runtime.NumCPU(),
	}
}

func (p *ParallelPrimeChecker) ParallelFindPrimesInRange(start, end int) []int {
	results := make([][]int, p.numWorkers)
	var wg sync.WaitGroup

	rangeSize := (end - start + 1) / p.numWorkers

	for i := 0; i < p.numWorkers; i++ {
		wg.Add(1)
		workerStart := start + (i * rangeSize)
		workerEnd := workerStart + rangeSize - 1
		if i == p.numWorkers-1 {
			workerEnd = end
		}

		go func(workerID, s, e int) {
			defer wg.Done()
			var localPrimes []int
			for n := s; n <= e; n++ {
				if IsPrime(n) {
					localPrimes = append(localPrimes, n)
				}
			}
			results[workerID] = localPrimes
		}(i, workerStart, workerEnd)
	}

	wg.Wait()

	var merged []int
	for _, r := range results {
		merged = append(merged, r...)
	}

	sort.Ints(merged)
	return merged
}
