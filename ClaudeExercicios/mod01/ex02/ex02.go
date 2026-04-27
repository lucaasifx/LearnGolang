package main

import "fmt"

type Permission uint
const (
	Read Permission = 1 << iota
	Write
	Execution
)

func main() {

	perm := Read | Write

	fmt.Printf("Leitura: %v\n", perm & Read != 0)
	fmt.Printf("Escrita: %v\n", perm & Write != 0)
	fmt.Printf("Execução: %v\n", perm & Execution != 0)

}