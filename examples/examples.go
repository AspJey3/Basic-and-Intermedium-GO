package examples

import (
	"fmt"
	"strings"
)

func esPalindromoSplit(palabra string) bool {
	palabra = strings.ToLower(palabra)
	letras := strings.Split(palabra, "")
	invertida := ""

	for i := len(letras) - 1; i >= 0; i-- {
		invertida += letras[i]
	}

	return palabra == invertida

}

func ExamplesCode() {

	fmt.Println(esPalindromoSplit("AnitaLavaLaTina")) // true
	fmt.Println(esPalindromoSplit("Golang"))          // false
}
