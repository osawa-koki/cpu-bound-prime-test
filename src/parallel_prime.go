package main

import (
	"runtime"
	"sort"
	"sync"
)

// ParallelPrimeChecker は並列処理による素数判定を行う構造体です
type ParallelPrimeChecker struct {
	numWorkers int
}

// NewParallelPrimeChecker は新しいParallelPrimeCheckerを作成します
func NewParallelPrimeChecker() *ParallelPrimeChecker {
	return &ParallelPrimeChecker{
		numWorkers: runtime.NumCPU(), // デフォルトはCPUコア数
	}
}

// SetWorkers はワーカー数を設定します
func (p *ParallelPrimeChecker) SetWorkers(n int) {
	if n > 0 {
		p.numWorkers = n
	}
}

// ParallelFindPrimesInRange は指定された範囲内の素数を並列に見つけます
func (p *ParallelPrimeChecker) ParallelFindPrimesInRange(start, end int) []int {
	// 範囲が小さい場合は逐次処理
	if end-start < 1000 {
		var primes []int
		for n := start; n <= end; n++ {
			if IsPrime(n) {
				primes = append(primes, n)
			}
		}
		return primes
	}

	// 各ワーカーの結果を保持するスライス
	results := make([][]int, p.numWorkers)
	var wg sync.WaitGroup

	// 範囲を分割
	rangeSize := (end - start + 1) / p.numWorkers
	if rangeSize < 1 {
		rangeSize = 1
	}

	// 各ワーカーを起動
	for i := 0; i < p.numWorkers; i++ {
		wg.Add(1)
		workerStart := start + (i * rangeSize)
		workerEnd := workerStart + rangeSize - 1
		if i == p.numWorkers-1 {
			workerEnd = end
		}

		go func(workerID, s, e int) {
			defer wg.Done()
			// ワーカーごとのローカルスライス
			localPrimes := make([]int, 0, (e-s+1)/2) // 容量を予め確保
			for n := s; n <= e; n++ {
				if IsPrime(n) {
					localPrimes = append(localPrimes, n)
				}
			}
			results[workerID] = localPrimes
		}(i, workerStart, workerEnd)
	}

	// すべてのワーカーの完了を待つ
	wg.Wait()

	// 結果をマージ
	totalSize := 0
	for _, r := range results {
		totalSize += len(r)
	}

	merged := make([]int, 0, totalSize)
	for _, r := range results {
		merged = append(merged, r...)
	}

	// ソート
	sort.Ints(merged)
	return merged
}
