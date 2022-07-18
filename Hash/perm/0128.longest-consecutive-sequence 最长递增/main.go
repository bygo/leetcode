package main

// https://leetcode.cn/problems/longest-consecutive-sequence

// ❓ 最长递增序列

func longestConsecutive(nums []int) int {
	numMpBool := map[int]bool{}
	for _, num := range nums {
		numMpBool[num] = true
	}

	var cntLongest int
	for num := range numMpBool {
		if !numMpBool[num-1] {
			cnt := 1
			num += 1
			for numMpBool[num] {
				num++
				cnt++
			}
			if cntLongest < cnt {
				cntLongest = cnt
			}
		}
	}
	return cntLongest
}
