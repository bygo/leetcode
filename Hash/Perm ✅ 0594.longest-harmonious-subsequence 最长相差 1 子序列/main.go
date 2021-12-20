package main

// https://leetcode-cn.com/problems/longest-harmonious-subsequence

// ❓ 最长相差1 子序列
// 二次扫描
func findLHSTwo(nums []int) int {
	numMpCnt := map[int]int{}
	for _, num := range nums {
		numMpCnt[num]++
	}

	var cntLongest int
	for num, cnt := range numMpCnt {
		if 0 < numMpCnt[num+1] {
			cntCur := numMpCnt[num+1] + cnt
			if cntLongest < cntCur {
				cntLongest = cntCur
			}
		}
	}
	return cntLongest
}

// 一次扫描
func findLHSOne(nums []int) int {
	numMpCnt := map[int]int{}
	var cntLongest int
	for _, num := range nums {
		numMpCnt[num]++
		if 0 < numMpCnt[num+1] {
			cntCur := numMpCnt[num+1] + numMpCnt[num]
			if cntLongest < cntCur {
				cntLongest = cntCur
			}
		}

		if 0 < numMpCnt[num-1] {
			cntCur := numMpCnt[num-1] + numMpCnt[num]
			if cntLongest < cntCur {
				cntLongest = cntCur
			}
		}
	}
	return cntLongest
}
