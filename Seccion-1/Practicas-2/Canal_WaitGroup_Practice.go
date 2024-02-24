package main

import (
	"fmt"
	"time"
	"sync"
	"math/rand/v2"
)
func getNumber(id int,  wg *sync.WaitGroup, canal *chan int) {
	defer wg.Done()
	seconds := rand.IntN(9) + 3
	fmt.Printf("Retrieving value: %d\n", id)
	time.Sleep(time.Duration(seconds) * time.Second)
	value := rand.IntN(60)
	fmt.Printf("Value %d retrieved: %d. Time elapsed: %d\n", id, value, seconds)
	*canal <- value
}
func main(){
	numOfValues := 9
	sumOfValues := 0

	var wg sync.WaitGroup
	wg.Add(numOfValues)

	fmt.Printf("Calculating sum of %d values\n", numOfValues)
	canal := make(chan int)

	for i := 0; i < numOfValues; i++ {
		go getNumber(i+1, &wg, &canal)
	}

	go func() {
		wg.Wait() // Esperamos a que todas las goroutines hayan terminado
		close(canal) // Cerramos el canal para indicar que ya no se enviarán más valores
	}()

	for result := range canal {
		sumOfValues += result
	}

	fmt.Printf("Sum completed: %d", sumOfValues)
}
