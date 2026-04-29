package main


import (
	"fmt"
	"slices"
)


func removeDuplicatas(s []int) []int {
	r := []int{}
	for _, value := range s {
		if !slices.Contains(r, value) {
			r = append(r, value)
		}
	}
	return r
}


func main() {
	// declarando um slice

	//var s []string
	//fmt.Println("uninit:", s, s == nil, len(s) == 0)
	// Retorna [] true true
		// zero value == nil
	//  s := make([]string, 3)
	//  fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))


	// Exercício 1 — fácil
	// Declare um slice com os números de 1 a 10 usando append num loop. Depois fatie e imprima apenas os números pares de posição (índices 1, 3, 5, 7, 9).
	// s := []int{}
	// for i := range 10 {
	// 	s = append(s, i+1)
	// }
	// for _, value := range s {
	// 	if value % 2 == 0 {
	// 		fmt.Println(value)
	// 	}
	// } 
	//Exercício 2 — médio
	//Dada essa função:
	// func removeDuplicatas(s []int) []int
	// Implemente ela sem usar maps — só com loops e slices. Entrada [1, 2, 2, 3, 3, 3, 4] deve retornar [1, 2, 3, 4].
	test := []int{1,2,2,3,3,3,4}
	test  = removeDuplicatas((test))
	fmt.Println(test)
	
}