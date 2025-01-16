package main

type SequentialPrimeChecker struct{}

func NewSequentialPrimeChecker() *SequentialPrimeChecker {
	return &SequentialPrimeChecker{}
}

func (s *SequentialPrimeChecker) SequentialIsPrime(n int) bool {
	return IsPrime(n)
}

func (s *SequentialPrimeChecker) SequentialFindPrimesInRange(start, end int) []int {
	var primes []int
	for i := start; i <= end; i++ {
		if s.SequentialIsPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}
