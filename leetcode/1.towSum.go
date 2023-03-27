package main

import (
	"fmt"
)

func main() {

	var useCases = []struct {
		nums   []int
		target int
	}{
		{
			nums:   []int{2, 7, 11, 15},
			target: 9,
		},
		{
			nums:   []int{3, 2, 4},
			target: 6,
		},
		{
			nums:   []int{3, 3},
			target: 6,
		},
	}

	for _, useCase := range useCases {
		fmt.Println(twoSum1(useCase.nums, useCase.target))
		fmt.Println(twoSum2(useCase.nums, useCase.target))
	}

}

func twoSum1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		num1 := nums[i]
		for j := i + 1; j < len(nums); j++ {
			if num1+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSum2(nums []int, target int) []int {
	vimap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		j, ok := vimap[target-nums[i]]
		if ok && j != i {
			return []int{i, j}
		}
		vimap[nums[i]] = i
	}

	return nil
}
