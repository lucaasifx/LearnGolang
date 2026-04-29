package main

import(
	"fmt"
	"math/rand/v2"
)


func main() {
	// Crie um programa que utilize a declaração switch, sem switch statement especificado.
	tomarCafe := rand.IntN(3)
	fmt.Println(tomarCafe)

	switch {
		case tomarCafe <= 0:
			fmt.Println("Hora de dormir.... zZzZzZ")
		case tomarCafe <= 1:
			fmt.Println("Que tal uma pausa para a hidratação?")
		case tomarCafe <= 2:
			fmt.Println("Hora do café :=")
		default:
			fmt.Println("Talvez vc curta mais cerveja :p")
	}
}