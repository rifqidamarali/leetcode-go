package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num1 := range nums {
		num2 := target - num1
		if index, found := numMap[num2]; found {
			return []int{index, i}
		} else {
			numMap[num1] = i
		}
	}

	return nil
}
func main() {
	nums := []int{2, 3, 1, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Println("Values:", result)
}