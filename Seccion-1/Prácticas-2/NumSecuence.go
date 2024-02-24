package main

import (
"fmt"
"math/rand"
)

type NumberSequence interface {
	Next() int
	First30() []int
	Title()
}

type IncreasingNumbers struct {
	numbers []int
}

type PrimeNumbers struct {
	numbers []int
}


func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func printArray(arr []int) {
	fmt.Print("[")
	for i, value := range arr {
		fmt.Print(value)
		if i < len(arr)-1 {
			fmt.Print(", ")
		}
	}
}

func main() {

}
