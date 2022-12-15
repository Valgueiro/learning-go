package main

import "fmt"

func main() {
	fmt.Print(romanToInt("IX"))
}

func contains[T comparable](s []T, e T) bool {
	for _, element := range s {
		if element == e {
			return true
		}
	}

	return false
}

// wihtout generics
// func contains(s []byte, e byte) bool {
// 	for _, element := range s {
// 		if element == e {
// 			return true
// 		}
// 	}

// 	return false
// }

func getNumberFromSymbol(b byte) int {
	symbolToNumber := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	return symbolToNumber[b]
}

func shouldSubtract(b byte, next_b byte) bool {
	subtractRules := map[byte][]byte{
		'I': {'V', 'X'},
		'X': {'L', 'C'},
		'C': {'D', 'M'},
	}

	return contains(subtractRules[b], next_b)
}

func romanToInt(s string) int {
	acumulator := 0
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && shouldSubtract(s[i], s[i+1]) {
			acumulator += getNumberFromSymbol(s[i+1]) - getNumberFromSymbol(s[i])
			i++
		} else {
			acumulator += getNumberFromSymbol(s[i])
		}

	}
	return acumulator
}
