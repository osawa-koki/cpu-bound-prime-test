package main

import (
	"fmt"
	"time"
)

func main() {
	start := 1
	end := 1000000

	{
		sequentialChecker := NewSequentialPrimeChecker()
		startTime := time.Now()
		primes := sequentialChecker.SequentialFindPrimesInRange(start, end)
		duration := time.Since(startTime)

		fmt.Println("----- ----- ----- ----- -----")
		fmt.Println("SequentialPrimeChecker")
		fmt.Println("From", start, "to", end)
		fmt.Println("Prime count:", len(primes))
		fmt.Println("Execution time:", duration)
		fmt.Println("----- ----- ----- ----- -----")
	}
}
