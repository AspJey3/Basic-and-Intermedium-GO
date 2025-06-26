package code

import (
	"fmt"
	"strings"
)

// multiplicador recibe un número y muestra su tabla de multiplicar del 0 al 10

func multiplicador(number int) {

	result := 0

	for i := 0; i <= 10; i++ {
		result = number * i
		fmt.Println(number, "x", i, "=", result)
	}
}

// nFibonacci recibe un número n y devuelve los n primeros números de la serie de Fibonacci
func nFibonacci(n int) []int {

	if n <= 0 {
		return []int{}
	}

	if n == 1 {
		return []int{0}
	}

	serie := []int{0, 1}

	for i := 2; i < n; i++ {
		serie = append(serie, serie[i-2]+serie[i-1])
	}

	return serie

}

// InvertirString recibe una palabra y devuelve la palabra invertida

func invertirString(word string) bool {
	reversedLowerCase := strings.ToLower(word)
	letters := strings.Split(reversedLowerCase, "")
	reversed := ""

	for i := len(letters) - 1; i >= 0; i-- {
		reversed += letters[i]
	}

	return word == reversed
}

func mergeAlternately(word1, word2 string) string {

	maxLength := len(word1)
	mergeLetters := ""

	for i := 0; i < len(word1); i++ {
		if i < maxLength {
			mergeLetters += string(word1[i])
		}
		if i < maxLength {
			mergeLetters += string(word2[i])
		}
	}

	return mergeLetters
}

func LeetCode() {

	// Ejercicio 1
	// const num int = 5
	// multiplicador(num)

	// Ejercicio 2

	// const word string = "Hello"

	// fmt.Println(InvertirString(word))

	// Ejercicio 3

	// const n int = 8
	// fmt.Println(nFibonacci(n))

	//Ejercicio 4
	// const word string = "Juan"
	// fmt.Println(invertirString(word))

	// Ejercicio 5
	const word1 string = "abc"
	const word2 string = "123"

	fmt.Println(mergeAlternately(word1, word2))

}
