package theory

import "fmt"

func Conditionals() {

	number := 4

	fmt.Println("--------------- Condicionales y Switch -----------------")
	// Estructura condicional if-else || else if
	if number%2 == 0 {
		fmt.Println("Es un numero par")
	} else {
		fmt.Println("Es un numero impart")
	}

	// Estructura condicional switch

	//EJEMPLO 1

	switch day := "martes"; day {
	case "lunes":
		fmt.Println("Inicio de la semana")
	case "martes":
		fmt.Println("Segundo día")
	default:
		fmt.Println("Otro día")
	}

	//EJEMPLO 2

	color := "verde"

	switch color {
	case "rojo":
		fmt.Println("El color es rojo")
	case "verde":
		fmt.Println("El color es verde")
	default:
		fmt.Println("El color es cualquiera menos rojo o verde")
	}

	fmt.Println("-------------------------------------------")
}
