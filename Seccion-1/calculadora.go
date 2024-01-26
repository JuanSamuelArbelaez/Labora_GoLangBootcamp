package main

import "fmt"

//Pair Programming - Calculadora de Área y Perímetro

// Adrian, Gastón, Juan

func main() {
	errMsg := "Error. Ingrese un número valido"

	width := readInt("Ingrese la longitud del rectangulo:", errMsg)
	height := readInt("Ingrese la altura del rectangulo:", errMsg)

	fmt.Println("El area del rectangulo es:", computeArea(height, width))
	fmt.Println("El perimetro del rectangulo es:", computePerimeter(height, width))
}

func computeArea(height, width int) int {
	return height*width
}

func computePerimeter(height, width int) int {
	return 2*(height+width)
}

func readInt(prompt string, errorMsg string) int {
	var num float64
	for {
		fmt.Print(prompt)
		_, err := fmt.Scan(&num)

		if err != nil{
			panic(errorMsg)
		} else {break}
	}
	return int(num)
}
