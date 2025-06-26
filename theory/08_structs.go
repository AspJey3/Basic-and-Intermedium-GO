package theory

import "fmt"

type User struct {
	Nombre   string
	Edad     int
	Email    string
	Password string
}

// Metodo asociado a la estructura User
func (u User) getNombre() string {
	return u.Nombre
}

func login(user User) {
	fmt.Println("Usuario registrado:")
	fmt.Println("Nombre:", user.Nombre)
	fmt.Println("Edad:", user.Edad)
	fmt.Println("Email:", user.Email)
	// No imprimimos la contraseña por seguridad
}

func register(nombre string, edad int, email string, password string) User {
	return User{
		Nombre: nombre,
		Edad:   edad,
		Email:  email,
		// No imprimimos la contraseña por seguridad
		// Pero para este caso practico vamos a imprimirla
		Password: password,
	}
}

func Structures() {
	user := User{Nombre: "Juan", Edad: 30, Email: "juan@example", Password: "1234"}
	fmt.Println("Nombre obtenido con getNombre():", user.getNombre())

	login(user)
	fmt.Println("Register:", register(user.Nombre, user.Edad, user.Email, user.Password))

	// ARRAY DE OBJECTOS (COMO SE DIRIA EN JS/TS)

	users := []User{
		{Nombre: "Juan", Edad: 30, Email: "juan@example.com", Password: "1234"},
		{Nombre: "Ana", Edad: 25, Email: "ana@example.com", Password: "abcd"},
		{Nombre: "Luis", Edad: 40, Email: "luis@example.com", Password: "xyz1"},
	}

	// IMPORTANTE: NO ES LO MISMO QUE DICCIONARIOS for key, value := range diccionario {} XXXX
	for i, user := range users {
		fmt.Println("Usuario", i+1)
		fmt.Println("Nombre:", user.Nombre)
		fmt.Println("Edad:", user.Edad)
		fmt.Println("Email:", user.Email)
		fmt.Println("Password:", user.Password)
		fmt.Println("--------------------")
	}

}
