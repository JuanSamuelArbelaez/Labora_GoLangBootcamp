package main

import (
	"fmt"
	"sort"
)

var planets= map[int]Planet{
	1: {"Mercury", 0},
	2: {"Venus", 0},
	3: {"Earth", 1},
	4: {"Mars", 2},
	5: {"Jupiter", 95},
	6: {"Saturn",146},
	7: {"Uranus",27},
	8: {"Neptune",14},
}

type Planet struct {
	Name string
	NumberOfMoons int
}

func main() {
	printAllPlanets()
}

func printPlanet(planetId int){
	if planet, ok := planets[planetId]; ok {
		fmt.Println(planet.Name, ":", planet.NumberOfMoons)
	} else {
		panic(fmt.Sprintf("No planet with id <%d>", planetId))
	}
}

func printAllPlanets(){
	var sortedIDs []int
	for planetId := range planets {
		sortedIDs = append(sortedIDs, planetId)
	}
	sort.Ints(sortedIDs)

	for _, planetId := range sortedIDs {
		printPlanet(planetId)
	}
}
