package main

// https://leetcode-cn.com/problems/maximum-number-of-balls-in-a-box

func countBalls(lowLimit int, highLimit int) int {
	var sum int
	j := lowLimit
	for 0 < j {
		sum += j % 10
		j /= 10
	}
	m := [46]int{} // 最大9999 = 46
	var max int
	for i := lowLimit; i <= highLimit; i++ {
		m[sum]++
		k := i
		for k%10 == 9 { // 每次进1 减9
			sum -= 9
			k /= 10
		}
		sum++
	}

	for i := range m {
		if max < m[i] {
			max = m[i]
		}
	}

	return max
}
