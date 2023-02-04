package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var result []int

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result = append(result, i)
				result = append(result, j)
			}
		}
	}

	return result
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	s := twoSum(nums, target)
	fmt.Println(s)
}
