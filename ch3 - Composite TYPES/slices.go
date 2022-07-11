package main

import "fmt"

func main() {
	// var x = []int{10, 20}
	// fmt.Println(x)

	// // specify on ly the indices:
	// var y = []int{1, 3, 5: 2, 6, 11: 29, 15}
	// fmt.Println(y) // [1 3 0 0 0 2 6 0 0 0 0 29 15]

	var x []int // has value of nil

	x = append(x, 10)      // x = [10]
	x = append(x, 1, 2, 3) // x = [10, 1, 2, 3]

	var y = []int{2, 5}
	x = append(x, y...) // x = [10 1 2 3 2 5]
	fmt.Println(x)

	w := make([]int, 5)
	fmt.Println(w)

	f := make([]int, 5, 10) // define the capacity right away
	fmt.Println(f, len(f), cap(f))
}
