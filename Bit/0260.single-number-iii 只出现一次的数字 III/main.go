package main

// https://leetcode-cn.com/problems/single-number-iii

func singleNumber(nums []int) []int {
	sum := 0
	for i := range nums {
		sum ^= nums[i]
	}

	sum = sum & -sum // 数字不同 肯定会有一位不同
	var t1, t2 int
	for i := range nums {
		if 0 < nums[i]&sum {
			t1 ^= nums[i]
		} else {
			t2 ^= nums[i]
		}
	}
	return []int{t1, t2}
}
