package main

// https://leetcode-cn.com/problems/maximum-xor-of-two-numbers-in-an-array

func findMaximumXOR(nums []int) int {
	var res int
	const h = 30
	for k := h; 0 <= k; k-- {
		next := res*2 + 1
		res = next - 1

		m := map[int]bool{}
		for _, num := range nums {
			m[num>>k] = true
		}

		for _, num := range nums {
			if m[num>>k^next] {
				res += 1
				break
			}
		}
	}
	return res
}
