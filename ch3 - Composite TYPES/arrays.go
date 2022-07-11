package main

import "fmt"

func main() {
	// var x [3]int // starts with 0
	// var x = [3]int{10, 20, 30}
	// var x = [...]int{10, 20, 30}

	var x = [...]int{1, 2, 3}
	var y = [3]int{1, 2, 3}
	fmt.Println(x == y) // it works!

	var z [2][3]int // multi dimensional

	x[0] = 10
	fmt.Println(x[2])

	length := len(x)
}
