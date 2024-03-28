package main

import (
	"fmt"
	"os"
)

// Function to find prime factors
func PrimeFactors(n int) []int {
	factors := []int{}
	// Divide by 2 until it's no longer divisible
	for n%2 == 0 {
		factors = append(factors, 2)
		n /= 2
	}
	// Check odd numbers for factors
	for i := 3; i*i <= n; i += 2 {
		for n%i == 0 {
			factors = append(factors, i)
			n /= i
		}
	}
	// If n is a prime greater than 2
	if n > 2 {
		factors = append(factors, n)
	}
	return factors
}

// Function to convert string to int
func atoi(s string) (int, error) {
	result := 0
	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			result = result*10 + int(ch-'0')
		} else {
			return 0, fmt.Errorf("strconv.Atoi: parsing %s: invalid syntax", s)
		}
	}
	return result, nil
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println()
		return
	}
	num, err := atoi(args[0])
	if err != nil || num < 1 {
		fmt.Println()
		return
	}
	factors := PrimeFactors(num)
	for i, factor := range factors {
		fmt.Print(factor)
		if i != len(factors)-1 {
			fmt.Print("*")
		}
	}
	fmt.Println()
}
