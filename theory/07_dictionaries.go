package theory

import "fmt"

func Dictionaries() {

	// SE DEBE INICIALIZAR CON DATA

	ages := map[string]int{
		"Juan":  20,
		"Pedro": 30,
		"Maria": 25,
	}
	fmt.Println("diccionario:", ages)

	// Recorrer un diccionario
	for name, age := range ages {
		fmt.Printf("Nombre: %s, Edad: %d\n", name, age)
	}

	// ESTE ES PARA CREAR UN DICCIONARIO VACÍO

	users := make(map[string]string)

	fmt.Println("diccionario con make:", users)
	// Agregar elementos al diccionario
	users["Adress"] = "Calle Falsa 123"
	users["Phone"] = "123-456-7890"

	//recorrer un diccionario
	for key, value := range users {
		fmt.Printf("Clave: %s, Valor: %s\n", key, value)
	}
}
