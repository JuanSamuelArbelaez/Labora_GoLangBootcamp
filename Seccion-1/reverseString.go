package main

import (
	"fmt"
)

func main() {
	msg := "Input a string: "
	str := read(msg)
	fmt.Println("Your reversed string is:", reverse(str))
}

func read(prompt string) string {
	var str string
	fmt.Print(prompt)
	_, err := fmt.Scan(&str)
	if err != nil {panic(err)}
	return str
}

func reverse(str string) string {
	runes := []rune(str)

	// Credit to user Yazu in Stackoverflow for rune refactor of conventional method
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
