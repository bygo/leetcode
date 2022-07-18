package main

// https://leetcode.cn/problems/maximize-sum-of-array-after-k-negations

func largestSumAfterKNegations(nums []int, k int) int {
	var res int
	freq := map[int]int{}
	for _, num := range nums {
		freq[num]++
		res += num
	}

	// 负数
	for i := -100; i < 0 && 0 < k; i++ {
		if 0 < freq[i] {
			cnt := min(k, freq[i])
			res -= i * cnt * 2
			freq[-i] += cnt // 现有的数
			k -= cnt
		}
	}

	// 正数
	// 奇数 没有0
	if 0 < k && k%2 == 1 && freq[0] == 0 {
		for i := 1; i <= 100; i++ {
			if 0 < freq[i] { // 找到第一个
				res -= i * 2
				break
			}
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
