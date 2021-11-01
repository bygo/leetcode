package main

// https://leetcode-cn.com/problems/maximum-number-of-balls-in-a-box

func countBalls(lowLimit int, highLimit int) int {
	var sum int
	j := lowLimit
	for 0 < j {
		sum += j % 10
		j /= 10
	}
	m := [46]int{}
	var max int
	for i := lowLimit; i <= highLimit; i++ {
		m[sum]++
		if max < m[sum] {
			max = m[sum]
		}
		k := i
		for k%10 == 9 { // 每次进1 减9
			sum -= 9
			k /= 10
		}
		sum++
	}

	return max
}
