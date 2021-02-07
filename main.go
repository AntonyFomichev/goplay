package main

import "fmt"

func twoSum(nums [4]int, target int) [2]int {
	var result [2]int

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums) ; j++ {
			if nums[i] + nums[j] == target {
				result[0] = i
				result[1] = j
			}
		}
	}

	return result
}

func main() {
	nums := [4]int{2,7,11,15}
	target := 9

	fmt.Println(twoSum(nums, target))
}