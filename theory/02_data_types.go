package theory

import "fmt"

func DataTypes() {
	var integer int = 10

	// usa 32 bits (4bytes). Tiene precisión de unos 6 - 7 dígitos decimales
	var floatNumber float32 = 10.5
	// usa 64 bits (8bytes). Tiene precisión de unos 15 - 16 dígitos decimales
	var floatNumber2 float64 = 10.4

	fmt.Println("--------------- Data Types -----------------")
	// -------------------------------------- EJEMPLO ----------------------------------- //
	var a float32 = 3.141592653589793
	var b float64 = 3.141592653589793
	fmt.Println("float32:", a) // Imprime: 3.1415927
	fmt.Println("float64:", b) // Imprime: 3.141592653589793
	// ---------------------------------------------------------------------------------- //

	// tipo booleano, puede ser true o false
	var boolean bool = true

	// tipo string, es una cadena de caracteres
	var stringExample string = "Golang es un lenguaje de programación"

	fmt.Println(integer, floatNumber, floatNumber2, boolean, stringExample)
	fmt.Println("-------------------------------------------")

}
