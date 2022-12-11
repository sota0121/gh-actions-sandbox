package main

import (
	"fmt"

	"github.com/sota0121/gh-actions-sandbox/lib"
)

func main() {
	fmt.Println("Hi, I'm the fizzbuzz CLI!")

	// Main loop
	for {
		fmt.Println("---\nPlease enter a number !")
		fmt.Println("(Enter 0 to exit)")
		var n int
		fmt.Scanln(&n)

		// Exit if n is 0
		if n == 0 {
			fmt.Println("Bye!")
			return
		}

		result, err := lib.FizzBuzz(n)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		fmt.Println()
	}
}
