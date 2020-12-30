package main

import (
	"fmt"
	"math"
)

func canJump(nums []int) bool {
	i := 0

	//var max_index int
	max_index := nums[i]
	for i < len(nums) && i <= max_index {
		new_index := nums[i] + i
		max_index = int(math.Max(float64(new_index), float64(max_index)))
		i += 1
	}
	if i == len(nums) {
		return true
	}
	return false
}
func main() {
	nums := []int{1, 3, 0, 0, 1, 4}
	if canJump(nums) {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
