package main

// https://leetcode-cn.com/problems/single-number-ii

func singleNumber(nums []int) int {
	a, b := 0, 0
	for _, num := range nums {
		b = (b ^ num) &^ a
		a = (a ^ num) &^ b
	}
	return b
}

func singleNumber(nums []int) int {
	res := int32(0)
	for i := 0; i < 32; i++ {
		total := int32(0)
		for _, num := range nums {
			total += int32(num) >> i & 1
		}
		if 0 < total%3 {
			res |= 1 << i
		}
	}
	return int(res)
}
