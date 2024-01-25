package main

import "fmt"

var planets= map[string]int{
	"Mercury":0,
	"Venus":0,
	"Earth":1,
	"Mars":2,
	"Jupiter":95,
	"Saturn":146,
	"Uranus":27,
	"Neptune":14,
}

func main() {
	printAllPlanets()
}

func printPlanet(planetName string){
	if numberOfMoons, ok := planets[planetName]; ok {
		fmt.Println(planetName, ":", numberOfMoons)
	} else {
		panic("No planet with name <" + planetName + ">")
	}
}

func printAllPlanets(){
	for planetName := range planets{
		printPlanet(planetName)
	}
}
