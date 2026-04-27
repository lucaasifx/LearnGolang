package main

import "fmt"


const (
	_ = 2026 + iota
	DAQUIA1
	DAQUIA2
	DAQUIA3
	DAQUIA4
)


func main() {
	// Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos.
	// Demonstre estes valores.

	fmt.Println(DAQUIA1)
	fmt.Println(DAQUIA2)
	fmt.Println(DAQUIA3)
	fmt.Println(DAQUIA4)

}