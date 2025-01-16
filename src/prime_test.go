package main

import (
	"testing"
)

func BenchmarkIsPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime(999983) // 大きな素数でテスト
	}
}

func BenchmarkFindPrimesInRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindPrimesInRange(1, 10000)
	}
}

func TestIsPrime(t *testing.T) {
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
			if got := IsPrime(tt.input); got != tt.expected {
				t.Errorf("IsPrime(%d) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}
