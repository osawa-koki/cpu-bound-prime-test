package main

import (
	"testing"
)

func BenchmarkSequentialFindPrimesInRange(b *testing.B) {
	checker := NewSequentialPrimeChecker()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		checker.SequentialFindPrimesInRange(1, 10000)
	}
}

func TestSequentialFindPrimesInRange(t *testing.T) {
	checker := NewSequentialPrimeChecker()
	tests := []struct {
		name     string
		start    int
		end      int
		expected int
	}{
		{"small range", 1, 100, 25},
		{"medium range", 1, 1000, 168},
		{"large range", 1000, 2000, 135},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			primes := checker.SequentialFindPrimesInRange(tt.start, tt.end)
			if got := len(primes); got != tt.expected {
				t.Errorf("SequentialFindPrimesInRange(%d, %d) returned %d primes, want %d",
					tt.start, tt.end, got, tt.expected)
			}
		})
	}
}
