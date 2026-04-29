package main

import (
	"math/rand/v2"
	"fmt"
)
func main() {
	// Utilizando a solução anterior, adicione as opções else if e else.
	
	// Randomiza um número inteiro entre 0 e 2
	tomarCafe := rand.IntN(3)
	fmt.Println(tomarCafe)
	if tomarCafe <= 0 {
		fmt.Println("Hora de dormir.... zZzZzZ")
	} else if tomarCafe <= 1 {
		fmt.Println("Que tal uma pausa para a hidratação?")
	} else {
		fmt.Println("Hora do café :=")
	}
}