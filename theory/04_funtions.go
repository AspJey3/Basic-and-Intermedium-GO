package theory

import "fmt"

// función que suma dos enteros

// PALABRA CLAVE: func
// sintexis: "func nombre(parametros) tipo_retorno { }"

func suma(a, b int) int {
	return a + b
}

// función con retorno múltiple
func divmod(a, b int) (int, int) {
	return a / b, a % b
}

// ATENCIÓN
// main es la función principal del programa
// func main() {

// }

func Functions() {
	total := suma(3, 5)
	fmt.Println("Suma =", total)

	q, r := divmod(10, 3)
	fmt.Printf("cociente = %d, resto = %d\n", q, r)
}
