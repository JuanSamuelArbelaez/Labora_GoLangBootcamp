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

// Función para mostrar las personas
func mostrarPersonas(personas []Persona) {
	for _, p := range personas {
		fmt.Printf("Nombre: %s, Edad: %d, Altura: %.2f, Peso: %.2f ", p.Nombre, p.Edad, p.Altura, p.Peso)
	}
}

func main(){

}
