package main

// https://leetcode-cn.com/problems/count-largest-group

// ❓ 数位和相同为一组，含有最大数量的有多少组

func countLargestGroup(n int) int {
	// 统计
	sumMpCnt := map[int]int{}
	for i := 1; i <= n; i++ {
		j := i
		sum := 0
		for 0 < j {
			sum += j % 10
			j /= 10
		}
		sumMpCnt[sum] ++
	}

	// 最大数量
	var cnt, cntLargest int
	for sum := range sumMpCnt {
		if cnt < sumMpCnt[sum] {
			cnt = sumMpCnt[sum]
			cntLargest = 1
		} else if sumMpCnt[sum] == cnt {
			cntLargest++
		}
	}
	return cntLargest
}
