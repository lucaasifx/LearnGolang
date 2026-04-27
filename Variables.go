// Inteiros — tamanho explícito
var a int8   = 127
var b int16  = 32767
var c int32  = 2147483647
var d int64  = 9223372036854775807
var e int    = 42       // tamanho varia com a arquitetura (32 ou 64 bits)
var u uint   = 42       // sem sinal

// Ponto flutuante
var f32 float32 = 3.14
var f64 float64 = 3.141592653589793  // padrão para literais decimais

// Booleano
var ok bool = true

// String — imutável, UTF-8 nativo
var msg string = "Olá, 世界"

// Rune — alias para int32, representa um caractere Unicode
var r rune = 'A'   // note aspas simples

// Byte — alias para uint8
var by byte = 255


// declaração completa
var x int = 10
// declaração com o descritor curto
x := 10
// AMBAS FUNCIONAM NORMALMENTE!


