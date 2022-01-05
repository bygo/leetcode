package main

// https://leetcode-cn.com/problems/maximum-xor-of-two-numbers-in-an-array

func findMaximumXOR(nums []int) int {
	var numRes, numMax int
	for k := 30; 0 <= k; k-- {
		// 结果翻倍
		numRes = numRes << 1
		// 下一个最大值
		numMax = numRes + 1

		numMpBool := map[int]bool{}
		for _, num := range nums {
			numMpBool[num>>k] = true
		}

		for _, numCur := range nums {
			numCur >>= k
			// 异或结合律
			num := numCur ^ numMax
			if numMpBool[num] {
				numRes += 1
				break
			}
		}
	}
	return numRes
}
