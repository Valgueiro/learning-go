package main

import "fmt"

func main() {
	fmt.Print(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}

func removeDuplicates(nums []int) []int {

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			nums = append(nums[:i-1], nums[i:]...)
			i--
		}
	}
	return nums
}
