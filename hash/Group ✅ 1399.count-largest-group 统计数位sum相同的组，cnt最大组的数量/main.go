package main

// https://leetcode-cn.com/problems/count-largest-group

// ❓ 统计数位sum相同的组，cnt最大组的数量

func countLargestGroup(n int) int {
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
	var cnt, res int
	for sum := range sumMpCnt {
		if cnt < sumMpCnt[sum] {
			cnt = sumMpCnt[sum]
			res = 1
		} else if sumMpCnt[sum] == cnt {
			res++
		}
	}
	return res
}
