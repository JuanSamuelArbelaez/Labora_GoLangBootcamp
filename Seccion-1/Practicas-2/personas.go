package main

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

		// Validación de datos
		if edad < 0 || altura <= 0 || peso <= 0 {
			panic("Error: Datos ingresados no válidos. La edad debe ser positiva, la altura debe ser mayor a cero y el peso debe ser mayor a cero.")
			i-- // Repetir la iteración
			continue
		}

		*personas = append(*personas, Persona{Nombre: nombre, Edad: edad, Altura: altura, Peso: peso})
	}
}

// Función para buscar una persona en el slice
func buscarPersona(personas []Persona, valorBuscado string) *Persona {
	for _, p := range personas {
		if p.Nombre == valorBuscado || fmt.Sprint(p.Edad) == valorBuscado ||
			fmt.Sprint(p.Altura) == valorBuscado || fmt.Sprint(p.Peso) == valorBuscado {
			return &p
		}
	}
	return nil
}

// Función para ordenar personas según el criterio elegido por el usuario
func ordenarPersonas(personas []Persona, criterio int) []Persona {
	switch criterio {
	case 1: // Ordenar por nombre
		sort.Slice(personas, func(i, j int) bool {
			return personas[i].Nombre < personas[j].Nombre
		})
	case 2: // Ordenar por edad
		sort.Slice(personas, func(i, j int) bool {
			return personas[i].Edad < personas[j].Edad
		})
	case 3: // Ordenar por altura
		sort.Slice(personas, func(i, j int) bool {
			return personas[i].Altura < personas[j].Altura
		})
	case 4: // Ordenar por peso
		sort.Slice(personas, func(i, j int) bool {
			return personas[i].Peso < personas[j].Peso
		})
	}
	return personas
}

func main(){
	var personas []Persona

	for {
		fmt.Println("======= MENÚ =======")
		fmt.Println("1. Registrar personas")
		fmt.Println("2. Buscar persona")
		fmt.Println("3. Listar personas")
		fmt.Println("4. Ordenar lista de personas")
		fmt.Println("5. Salir")
		fmt.Print("Ingrese la opción: ")

		var opcion int
		fmt.Scan(&opcion)

		switch opcion {
		case 1:
			// Registro de personas
			registrarPersonas(&personas)
		case 2:
			// Búsqueda de persona
			fmt.Print("Ingrese el valor a buscar: ")
			var valorBuscado string
			fmt.Scan(&valorBuscado)

			personaEncontrada := buscarPersona(personas, valorBuscado)
			if personaEncontrada != nil {
				fmt.Printf("Persona encontrada: Nombre: %s, Edad: %d, Altura: %.2f, Peso: %.2f\n", personaEncontrada.Nombre, personaEncontrada.Edad, personaEncontrada.Altura, personaEncontrada.Peso)
			} else {
				fmt.Println("Persona no encontrada.")
			}
		case 3:
			// Mostrar personas sin ordenar
			fmt.Println("Personas sin ordenar:")
			mostrarPersonas(personas)
		case 4:
			// Ordenar lista de personas
			fmt.Println("Seleccione criterio de ordenamiento:")
			fmt.Println("1. Por nombre")
			fmt.Println("2. Por edad")
			fmt.Println("3. Por altura")
			fmt.Println("4. Por peso")
			fmt.Print("Ingrese la opción: ")

			var criterioOrdenamiento int
			fmt.Scan(&criterioOrdenamiento)

			personasOrdenadas := ordenarPersonas(personas, criterioOrdenamiento)

			fmt.Println("Personas ordenadas:")
			mostrarPersonas(personasOrdenadas)
		case 5:
			// Salir del programa
			fmt.Println("¡Hasta luego!")
			return

		default:
			fmt.Println("Opción no válida. Por favor, ingrese una opción válida.")
		}
	}
}
