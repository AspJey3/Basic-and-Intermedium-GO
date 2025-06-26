package theory

import "fmt"

func Loops() {
	fmt.Println("--------------- Loops -----------------")

	// Bucle for
	for i := 0; i < 5; i++ {
		fmt.Println("Iteración:", i)
	}

	// Bucle while (usando for)
	j := 0
	for j < 5 {
		fmt.Println("Iteración while:", j)
		j++
	}

	// for infinito (break o return para salir)
	k := 0
	for {
		if k == 2 {
			break
		}
		fmt.Println("k =", k)
		k++
	}

	fmt.Println("-------------------------------------------")
}
