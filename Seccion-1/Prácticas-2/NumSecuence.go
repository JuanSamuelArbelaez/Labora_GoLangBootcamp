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

func (inc *IncreasingNumbers) Next() int {
	nextNumber := len(inc.numbers) + 1
	inc.numbers = append(inc.numbers, nextNumber)
	return nextNumber
}

func (prime *PrimeNumbers) Next() int {
	if prime.numbers == nil {
		prime.numbers = append(prime.numbers, 2) // Start with 2, the first prime
		return 2
	}

	lastNumber := prime.numbers[len(prime.numbers)-1]
	newNumber := lastNumber + 1
	for {
		if isPrime(newNumber) {
			prime.numbers = append(prime.numbers, newNumber)
			break
		}
		newNumber++
	}
	return newNumber
}

func (inc *IncreasingNumbers) First30() []int {
	for len(inc.numbers) < 30 {
		inc.Next()
	}
	return inc.numbers
}

func (prime *PrimeNumbers) First30() []int {
	for len(prime.numbers) < 30 {
		prime.Next()
	}
	return prime.numbers
}

func (inc *IncreasingNumbers) Title() {
	fmt.Println("Sequence of Increasing Numbers")
}

func (prime *PrimeNumbers) Title() {
	fmt.Println("Sequence of Prime Numbers")
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
	fmt.Println("]")
}

func main() {

}
