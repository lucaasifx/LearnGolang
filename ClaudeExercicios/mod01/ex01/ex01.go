package main

import (
	"fmt"
	"math"
)

type Kg float64
type Metros float64

func main() {
	var peso Kg = 70.0
	var altura Metros = 1.75

	var IMC float64 = float64(peso) / float64(math.Pow(float64(altura), 2))

	fmt.Printf("IMC = %v\n", IMC)


}