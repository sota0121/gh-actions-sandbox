package lib

import (
	"fmt"
)

var (
	fizz     = "Fizz"
	buzz     = "Buzz"
	fizzbuzz = "FizzBuzz"
)

// FizzBuzz returns a string based on the following rules:
func FizzBuzz(n int) (string, error) {
	// Check if n is greater than 0
	if n < 0 {
		return "", fmt.Errorf("n must be greater than 0")
	}

	// Check if n is divisible by 3 and 5
	if n%3 == 0 && n%5 == 0 {
		return fizzbuzz, nil
	}
	// Check if n is divisible by 3
	if n%3 == 0 {
		return fizz, nil
	}
	// Check if n is divisible by 5
	if n%5 == 0 {
		return buzz, nil
	}
	// Return n as a string
	return fmt.Sprintf("%d", n), nil
}
