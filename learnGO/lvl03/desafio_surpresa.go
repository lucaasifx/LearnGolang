package main

import "fmt"



func main() {
	// 	Desafio surpresa!
	// Format printing:
	// Decimal       %d
	// Hexadecimal   %#x
	// Unicode       %#U
	// Tab           \t
	// Linha nova    \n
	// Faça um loop dos números 33 a 122, e utilize format printing para demonstrá-los como texto/string.


	for x := 33; x <= 122; x++ {
		fmt.Printf("%v => %#U\n", x, x)
	}
}