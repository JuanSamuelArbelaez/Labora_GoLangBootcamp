package main

import "fmt"

func main() {
	var planetsMoons = map[string]int{
		"Mercury":0,
		"Venus":0,
		"Earth":1,
		"Mars":2,
		"Jupiter":95,
		"Saturn":146,
		"Uranus":27,
		"Neptune":14,
	}

	printPlanet(planetsMoons, "Mercury")
	printPlanet(planetsMoons, "Venus")
	printPlanet(planetsMoons, "Earth")
	printPlanet(planetsMoons, "Mars")
	printPlanet(planetsMoons, "Jupiter")
	printPlanet(planetsMoons, "Saturn")
	printPlanet(planetsMoons, "Uranus")
	printPlanet(planetsMoons, "Neptune")
}

func printPlanet(planets map[string]int, planetName string){
	if numberOfMoons, ok := planets[planetName]; ok {
		fmt.Println(planetName, ":", numberOfMoons)
	} else {
		panic("No planet with name <" + planetName + ">")
	}
}
