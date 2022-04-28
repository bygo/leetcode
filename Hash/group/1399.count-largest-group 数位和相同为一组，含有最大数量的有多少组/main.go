package main

// https://leetcode-cn.com/problems/count-largest-group

// ❓ 数位和相同为一组，含有最大数量的有多少组

func countLargestGroup(n int) int {
	// 计算初始数位和
	var sum = 1

	// 统计
	sumMpCnt := [46]int{} // 最大99999 = 46
	for num := 1; num <= n; num++ {
		sumMpCnt[sum] ++
		numCur := num
		for numCur%10 == 9 {
			sum -= 9
			numCur /= 10
		}
		sum++
	}

	// 最大数量
	var cntMax, cntLargest int
	for sum := range sumMpCnt {
		if cntMax < sumMpCnt[sum] {
			cntMax = sumMpCnt[sum]
			cntLargest = 1
		} else if sumMpCnt[sum] == cntMax {
			cntLargest++
		}
	}
	return cntLargest
}
