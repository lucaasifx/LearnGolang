package main

import (
	"fmt"
)

func main() {

	// Utilizando o operador curto de declaração, atribua estes valores às variáveis com os identificadores "x", "y", e "z".
	// 	1. 42
	// 	2. "James Bond"
	// 	3. true
	// Agora demonstre os valores nestas variáveis, com:
	// 	1. Uma única declaração print
	// 	2. Múltiplas declarações print

	x, y, z := 42, "James Bond", true


	// Com um
	fmt.Printf("Valores de x,y e z: [%v, %v, %v]\n\n", x, y, z)

	// Com vários
	fmt.Println("Valor de x: ", x)
	fmt.Println("Valor de y: ", y)
	fmt.Println("Valor de z: ", z)

}