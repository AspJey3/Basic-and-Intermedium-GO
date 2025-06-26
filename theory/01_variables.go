package theory

import "fmt"

func PrintVariables() {
	// Una declaración de variable con inferencia de tipo
	nameBrother := "Juansito"
	edad := 30

	// Esta es su declaración explícita
	var saludo string = "Hola, mi nombre es Pablo"

	fmt.Println("--------------- Variables -----------------")
	fmt.Println(saludo, "y tengo", edad, "años. Mi hermano es", nameBrother)
	fmt.Println("-------------------------------------------")
}
