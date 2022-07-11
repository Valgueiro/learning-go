package main

import "fmt"

func main() {
	var flag bool
	flag = true

	fmt.Println(flag)

	var isAwesome = true // infer type automatically
	fmt.Println(isAwesome)

	var str = "hello"

	str += ", Mateus!"
	fmt.Println(str)

	var intX int = 10
	var floatY float64 = 2.3

	// fmt.Println(intX + floatY) # it does not work
	fmt.Println(intX + int(floatY))
	fmt.Println(float64(intX) + floatY)

	short := 3.14

	fmt.Println(short)

	const y = "hello"

	// y = "ola" # it breaks
}
