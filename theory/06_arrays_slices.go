package theory

import "fmt"

func ArraySlices() {
	// array de longitud fija
	var a [3]int = [3]int{1, 2, 3}
	fmt.Println("array:", a)

	// slice (dinámico)
	s := []int{4, 5, 6}
	fmt.Println("slice inicial:", s)

	// agregar elementos
	s = append(s, 7, 8)
	fmt.Println("slice modificado:", s)

	// sub-slice
	subs := s[1:4]
	fmt.Println("sub-slice:", subs)

	fruits := [3]string{"manzana", "banana", "cereza"}
	sliceFruits := fruits[0:2]
	fmt.Println("fruits:", sliceFruits)

	//output: // fruits: [manzana banana]

	//-----------------------------------------------------------------------------------------------//

	// CASOS ESPECIALES BITS

	// TRANSFORMAR un string a bits
	message := "Hola"

	for i := 0; i < len(message); i++ {
		letra := message[i]
		fmt.Printf("Letra: %c | ASCII: %d | Binario: %08b\n", letra, letra, letra)
	}

	// TRANSFORMAR bits a string
	bits := []byte{72, 111, 108, 97}
	texto := string(bits)
	fmt.Println("Texto:", texto) // Texto: Hola

}
