package main

import (
	"fmt"
	"time"
)

// o ano de nascimento do leitor
const anoNasc = 2001

func main() {
	// Crie um loop utilizando a sintaxe: for condition {}
	// Utilize-o para demonstrar os anos desde que você nasceu.
	year := anoNasc
	for  year <= time.Now().Year() {
		fmt.Println(year)
		year++
	}
}