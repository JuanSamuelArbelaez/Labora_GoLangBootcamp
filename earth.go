package main

import "fmt"

func main() {
	var planetsMoons = make(map[string]any)

	planetsMoons["Mercury"] = "Unknown"
	planetsMoons["Venus"] = "Unknown"
	planetsMoons["Earth"] = 1
	planetsMoons["Mars"] = 2

	var planetName = "Mercury"

	if numberOfMoons, ok := planetsMoons[planetName]; ok {
		fmt.Println(planetName, ":", numberOfMoons)
	} else {
		panic("No planet with name <" + planetName + ">")
	}
}
