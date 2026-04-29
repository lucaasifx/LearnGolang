package main

import "fmt"

var x [5]int
func main() {

	fmt.Println(x)
	x[4] = 42
	fmt.Println(x)

	// utilizando o declarador curto
	y := [...]int{1, 2, 3, 4, 5}
	fmt.Println(y)

	// autocomplete com zero values
	b := [...]int{100, 2: 400, 500}
    fmt.Println("idx:", b)

	// percorrendo matrizes
	// example
	var twoD [2][3]int
    for i := range 2 {
        for j := range 3 {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)


	// apenas aplicando os conceitos
	var threeD[3][2][2]int
	for i := range 3 {
		for j := range 2 {
			for k := range 2 {
				threeD[i][j][k] = i + j + k
			}
		}
	}
	fmt.Println(threeD)

	// uma matriz zerada
	var exampleMatrix [2][4][4]int
	fmt.Println(exampleMatrix)

}