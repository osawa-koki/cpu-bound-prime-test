package main

import (
	"sort"
	"testing"
)

func BenchmarkParallelFindPrimesInRange(b *testing.B) {
	checker := NewParallelPrimeChecker()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		checker.ParallelFindPrimesInRange(1, 10000)
	}
}

func TestParallelFindPrimesInRange(t *testing.T) {
	checker := NewParallelPrimeChecker()
	sequential := NewSequentialPrimeChecker()

	// テストケース
	tests := []struct {
		name  string
		start int
		end   int
	}{
		{"small range", 1, 100},
		{"medium range", 1, 1000},
		{"large range", 1000, 2000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 並列処理での結果
			parallelPrimes := checker.ParallelFindPrimesInRange(tt.start, tt.end)
			// 逐次処理での結果
			sequentialPrimes := sequential.SequentialFindPrimesInRange(tt.start, tt.end)

			// 結果の長さを比較
			if len(parallelPrimes) != len(sequentialPrimes) {
				t.Errorf("ParallelFindPrimesInRange(%d, %d) returned %d primes, want %d",
					tt.start, tt.end, len(parallelPrimes), len(sequentialPrimes))
			}

			// 結果の内容を比較
			if !isEqualSlices(parallelPrimes, sequentialPrimes) {
				t.Errorf("ParallelFindPrimesInRange(%d, %d) returned different primes than sequential version",
					tt.start, tt.end)
			}
		})
	}
}

// スライスの比較ヘルパー関数
func isEqualSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	sort.Ints(a)
	sort.Ints(b)
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
