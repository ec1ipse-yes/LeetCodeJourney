package main

import "sort"

// Solution1: hash table
func containsDuplicate(nums []int) bool {
	nums_map := make(map[int]int, len(nums))
	for _, v := range nums {
		if _, ok := nums_map[v]; ok {
			return true
		} else {
			nums_map[v] = 1
		}
	}
	return false
}

// Solution2: sorting
func containsDuplicate2(nums []int) bool {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}
