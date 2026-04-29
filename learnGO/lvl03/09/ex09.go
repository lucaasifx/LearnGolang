package main

import (
	"fmt"
)


func main() {
	// Crie um programa que utilize a declaração switch, onde o switch 
	// statement seja uma variável do tipo string com identificador "esporteFavorito".
	esporteFavorito := "xadrez"
	switch  {
		case esporteFavorito == "xadrez":
			fmt.Println("Você é capivara rapaz!")
		case esporteFavorito == "futebol", esporteFavorito == "basquete":
			fmt.Println("Continue acompanhando seu time favorito!")
		default:
			fmt.Println("Não conheço esse esporte") 
	}

}