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
	rangeSize := (end - start + 1) / p.numWorkers

	var mu sync.Mutex
	var allPrimes []int
	var wg sync.WaitGroup

	for i := 0; i < p.numWorkers; i++ {
		wg.Add(1)
		workerStart := start + (i * rangeSize)
		workerEnd := workerStart + rangeSize - 1
		if i == p.numWorkers-1 {
			workerEnd = end
		}

		go func(s, e int) {
			defer wg.Done()
			var localPrimes []int

			for n := s; n <= e; n++ {
				if IsPrime(n) {
					localPrimes = append(localPrimes, n)
				}
			}

			mu.Lock()
			allPrimes = append(allPrimes, localPrimes...)
			mu.Unlock()
		}(workerStart, workerEnd)
	}

	wg.Wait()
	sort.Ints(allPrimes)
	return allPrimes
}
