package main

import (
	"sort"
	"testing"
)

func BenchmarkGoroutineFindPrimesInRange(b *testing.B) {
	checker := NewGoroutinePrimeChecker()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		checker.GoroutineFindPrimesInRange(1, 10000)
	}
}

func TestGoroutineFindPrimesInRange(t *testing.T) {
	checker := NewGoroutinePrimeChecker()
	sequential := NewSequentialPrimeChecker()

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
			goroutinePrimes := checker.GoroutineFindPrimesInRange(tt.start, tt.end)
			sequentialPrimes := sequential.SequentialFindPrimesInRange(tt.start, tt.end)

			if len(goroutinePrimes) != len(sequentialPrimes) {
				t.Errorf("GoroutineFindPrimesInRange(%d, %d) returned %d primes, want %d",
					tt.start, tt.end, len(goroutinePrimes), len(sequentialPrimes))
			}

			if !isEqualSlices(goroutinePrimes, sequentialPrimes) {
				t.Errorf("GoroutineFindPrimesInRange(%d, %d) returned different primes than sequential version",
					tt.start, tt.end)
			}
		})
	}
}

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
