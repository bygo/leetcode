package main

// https://leetcode-cn.com/problems/maximum-number-of-balls-in-a-box

// ❓ 盒子中小球的最大数量

func countBalls(lowLimit int, highLimit int) int {
	// 计算初始数位和
	var sum int
	numCur := lowLimit
	for 0 < numCur {
		sum += numCur % 10
		numCur /= 10
	}

	// 统计
	sumMpCnt := [46]int{} // 最大99999 = 46
	var cntMax int
	for num := lowLimit; num <= highLimit; num++ {
		sumMpCnt[sum]++
		numCur = num
		for numCur%10 == 9 { // 每次进1 减9
			sum -= 9
			numCur /= 10
		}
		sum++
	}

	// 最大数量
	for sum := range sumMpCnt {
		if cntMax < sumMpCnt[sum] {
			cntMax = sumMpCnt[sum]
		}
	}

	return cntMax
}
