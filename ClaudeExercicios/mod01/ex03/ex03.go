package main
import "fmt"

const (
    _  = iota // descartamos o primeiro valor
    KB = 1 << (10 * iota)
    MB // 2^20
    GB // 2^30
    TB // 2^40
)

func main() {
	// 2.5GB
    var tamanho uint64 = 2 * GB + 512 * MB

	fmt.Println("tamanho = ", tamanho << (10 * 1))
    // imprima em KB e MB
    //fmt.Println("Em KB:", tamanho)
    //fmt.Println("Em MB:", ???)
}