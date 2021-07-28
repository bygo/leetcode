package main

// https://leetcode-cn.com/problems/how-many-numbers-are-smaller-than-the-current-number/

func smallerNumbersThanCurrent(nums []int) []int {
	var cnt [101]int
	res := make([]int, len(nums))
	for _, num := range nums {
		cnt[num]++
	}

	for i := 1; i < 101; i++ {
		cnt[i] += cnt[i-1]
	}
	for i, num := range nums {
		if num != 0 {
			res[i] = cnt[num-1]
		}
	}
	return res
}
