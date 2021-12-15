package main

// https://leetcode-cn.com/problems/maximum-number-of-balls-in-a-box

// ❓ 盒子中小球的最大数量

func countBalls(lowLimit int, highLimit int) int {
	var sum int
	j := lowLimit
	for 0 < j {
		sum += j % 10
		j /= 10
	}
	sumMpCnt := [46]int{} // 最大9999 = 46
	var max int
	for i := lowLimit; i <= highLimit; i++ {
		sumMpCnt[sum]++
		k := i
		for k%10 == 9 { // 每次进1 减9
			sum -= 9
			k /= 10
		}
		sum++
	}

	for sum := range sumMpCnt {
		if max < sumMpCnt[sum] {
			max = sumMpCnt[sum]
		}
	}

	return max
}
