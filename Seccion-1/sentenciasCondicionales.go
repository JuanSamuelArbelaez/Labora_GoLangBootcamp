package main

import "fmt"

const (
	m=60
	h=3600
	d=86400
)

func main() {
	formatTime(1030)
	formatTime(12045)
	formatTime(176520)

	fmt.Println(areFriendPair(220, 284))
}

func formatTime(seg int){
	days := secondsToDays(seg)
	seg = seg - (days*d)

	hours := secondsToHours(seg)
	seg = seg - (hours*h)

	minutes := secondsToMinutes(seg)
	seg = seg - (minutes*m)

	fmt.Println("days=", days, "hours=", hours, "minutes=",  minutes, "seconds=",  seg)
}

func secondsToDays(seconds int) int {
	return seconds/d
}

func secondsToHours(seconds int) int {
	return seconds/h
}

func secondsToMinutes(seconds int) int {
	return seconds/m
}

func areFriendPair(a, b int) bool {
	divisorsA, divisorsB := getProperDivisors(a), getProperDivisors(b)
	sumDivA, sumDivB := sumIntArray(divisorsA), sumIntArray(divisorsB)
	return (sumDivA == b) && (sumDivB == a)
}

func getProperDivisors(num int) []int {
	var divisors []int
	for i :=1; i < num-1; i++ {
		if num%i == 0 {
			divisors = append(divisors,i)
		}
	}
	return divisors
}

func sumIntArray(nums []int) int {
	sum := 0

	for i :=0 ; i < len(nums); i++ {
		sum += nums[i]
	}

	return sum
}
