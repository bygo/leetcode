package main

// https://leetcode-cn.com/problems/contains-duplicate-ii

func containsNearbyDuplicate(nums []int, k int) bool {
	var set = map[int]int{}
	for i, num := range nums {
		v, ok := set[num]
		if ok {
			if i-v <= k {
				return true
			}
		}
		set[num] = i
	}

	return false
}


