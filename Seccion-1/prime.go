package main

import (
	"fmt"
	"math"
)

func main() {
	msg := "Enter a whole number greater than 1: "
	errorMsg := "Invalid input. Please enter a valid number."

	result := readWholeNumber(msg, errorMsg)
	fmt.Println(result, "is prime:", isPrime(result))
}

func readWholeNumber(prompt string, errorMsg string) int {
	var num float64
	for {
		fmt.Print(prompt)
		_, err := fmt.Scan(&num)

		if err != nil || num <= 1 || num != float64(int(num)) {
			panic(errorMsg)
		} else { break }
	}
	return int(num)
}

func isPrime(num int) bool {
	if num <= 1{ panic("Invalid number")}
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 { return false }
	}
	return true
}
