package Prácticas_2

import (
"fmt"
"math"
"sort"
)

// Definición de la estructura Persona
type Persona struct {
	Nombre string
	Edad   int
	Altura float64
	Peso   float64
}

// Función para calcular el índice de masa corporal (IMC)
func calcularIMC(persona Persona) float64 {
	return persona.Peso / math.Pow(persona.Altura, 2)
}

// Función para mostrar las personas con su IMC y categoría
func mostrarPersonas(personas []Persona) {
	for _, p := range personas {
		imc := calcularIMC(p)
		fmt.Printf("Nombre: %s, Edad: %d, Altura: %.2f, Peso: %.2f, IMC: %.2f - ", p.Nombre, p.Edad, p.Altura, p.Peso, imc)
		switch {
		case imc < 18.5:
			fmt.Println("Bajo peso")
		case imc >= 18.5 && imc <= 24.9:
			fmt.Println("Peso normal")
		case imc >= 25 && imc <= 29.9:
			fmt.Println("Sobrepeso")
		default:
			fmt.Println("Obesidad")
		}
	}
}

func main(){

}
