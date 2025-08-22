package main

import (
	"fmt"
	"slices"
	"strings"
	"time"
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

// Se te dan dos cadenas de texto, word1 y word2. Debes mezclar las cadenas alternando las letras, comenzando con word1.
// Si una de las cadenas es más larga que la otra, debes añadir sus letras restantes al final de la cadena mezclada.

func mergeAlternately(word1, word2 string) string {

	lendWord1 := len(word1)
	lendWord2 := len(word2)

	mergedWord := ""

	var maxLength int = lendWord1

	if lendWord1 < lendWord2 {
		maxLength = lendWord2
	}

	for i := 0; i < maxLength; i++ {
		if i < lendWord1 {
			mergedWord += string(word1[i])
		}

		if i < lendWord2 {
			mergedWord += string(word2[i])
		}
	}

	return mergedWord
}

func gcdOfStrings(str1, str2 string) string {

	if str1+str2 != str2+str1 {
		return ""
	}

	var gcd func(a, b int) int
	gcd = func(a, b int) int {
		if b == 0 {
			return a
		}
		return gcd(b, a%b)
	}

	gcdLen := gcd(len(str1), len(str2))

	return str1[:gcdLen]
}

func kidsWithCandies(candies []int, extraCandies int) []bool {

	var maxCandies int = 0
	newArrayBool := []bool{}

	for i := 0; i < len(candies); i++ {
		if candies[i] > maxCandies {
			maxCandies = candies[i]
		}
	}

	for i := 0; i < len(candies); i++ {
		candies[i] += extraCandies
		if candies[i] >= maxCandies {
			newArrayBool = append(newArrayBool, true)
		} else {
			newArrayBool = append(newArrayBool, false)
		}
	}

	return newArrayBool
}

// func canPlaceFlowers(flowerbed []int, n int) bool {

// 	// var output bool = false

// 	// for i:=0; i<len(flowerbed); i++{

// 	// 	if(flowerbed[i] == 0){

// 	// 	}

// 	// }

// 	return false
// }

func reverseVowels(s string) string {
	vowels := "aeiouAEIOU"
	splitString := strings.Split(s, "")

	reverseVowels := []string{}

	for i := 0; i < len(splitString); i++ {

		if strings.Contains(vowels, splitString[i]) {
			reverseVowels = append(reverseVowels, splitString[i])
		}

	}

	slices.Reverse(reverseVowels)
	vowelIndex := 0

	for i := 0; i < len(splitString); i++ {
		if strings.Contains(vowels, splitString[i]) {
			splitString[i] = reverseVowels[vowelIndex]
			vowelIndex++
		}

	}

	return strings.Join(splitString, "")
}

func reverseWords(s string) string {

	stringSplit := strings.Split(s, " ")

	newArray := []string{}

	for i := 0; i < len(stringSplit); i++ {
		if stringSplit[i] != "" {
			newArray = append(newArray, stringSplit[i])
		}
	}

	slices.Reverse(newArray)

	return strings.Join(newArray, " ")
}

func main() {
	start := time.Now()

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
	// const word1 string = "abcdf"
	// const word2 string = "123"

	// fmt.Println(mergeAlternately(word1, word2))

	// const word1 string = "AFGAFGAFG"
	// const word2 string = "AFG"

	// fmt.Println(word1 + word2)

	// fmt.Println(gcdOfStrings(word1, word2))

	// candies := []int{2, 8, 7}
	// const extraCandies int = 1

	// fmt.Println(kidsWithCandies(candies, extraCandies))

	// flowerbed := []int{1, 0, 0, 0, 1}
	// const n = 1

	// fmt.Println(canPlaceFlowers(flowerbed, n))

	// const s = "IceCreAm"

	// fmt.Println(reverseVowels(s))

	const s = "the  sky  is blue"
	fmt.Println(reverseWords(s))

	fmt.Printf("Tiempo: %v ms\n", time.Since(start).Milliseconds())
}
