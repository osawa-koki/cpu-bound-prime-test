package main

import (
	"testing"
)

func BenchmarkSequentialIsPrime(b *testing.B) {
	checker := NewSequentialPrimeChecker()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		checker.SequentialIsPrime(999983)
	}
}

func BenchmarkSequentialFindPrimesInRange(b *testing.B) {
	checker := NewSequentialPrimeChecker()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		checker.SequentialFindPrimesInRange(1, 10000)
	}
}

func TestSequentialIsPrime(t *testing.T) {
	checker := NewSequentialPrimeChecker()
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{"0 is not prime", 0, false},
		{"1 is not prime", 1, false},
		{"2 is prime", 2, true},
		{"3 is prime", 3, true},
		{"4 is not prime", 4, false},
		{"17 is prime", 17, true},
		{"25 is not prime", 25, false},
		{"997 is prime", 997, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checker.SequentialIsPrime(tt.input); got != tt.expected {
				t.Errorf("SequentialPrimeChecker.SequentialIsPrime(%d) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}
