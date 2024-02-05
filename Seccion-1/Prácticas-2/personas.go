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

//Función para registrar personas
func registrarPersonas(personas *[]Persona){
	for i := 0; i < 5; i++ {
		var nombre string
		var edad int
		var altura, peso float64

		fmt.Printf("Ingrese nombre de persona %d: ", i+1)
		fmt.Scan(&nombre)
		fmt.Printf("Ingrese edad de persona %d: ", i+1)
		fmt.Scan(&edad)
		fmt.Printf("Ingrese altura de persona %d: ", i+1)
		fmt.Scan(&altura)
		fmt.Printf("Ingrese peso de persona %d: ", i+1)
		fmt.Scan(&peso)

		*personas = append(*personas, Persona{Nombre: nombre, Edad: edad, Altura: altura, Peso: peso})
	}
}

func main(){
	var personas []Persona

	for {
		fmt.Println("======= MENÚ =======")
		fmt.Println("1. Registrar personas")
		fmt.Println("5. Salir")
		fmt.Print("Ingrese la opción: ")

		var opcion int
		fmt.Scan(&opcion)

		switch opcion {
		case 1:
			// Registro de personas
			registrarPersonas(&personas)
		case 5:
			// Salir del programa
			fmt.Println("¡Hasta luego!")
			return

		default:
			fmt.Println("Opción no válida. Por favor, ingrese una opción válida.")
		}
	}
}
