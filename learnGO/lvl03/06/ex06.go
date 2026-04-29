package main

import (
	"math/rand/v2"
	"fmt"
)
func main() {
	// Crie um programa que demonstre o funcionamento da declaração if.
	
	// Randomiza um número inteiro entre 0 e 1
	tomarCafe := rand.IntN(2)
	fmt.Println(tomarCafe)
	if tomarCafe <= 0 {
		fmt.Println("Hora de dormir.... zZzZzZ")
	} else {
		fmt.Println("Hora do café :p")
	}
}