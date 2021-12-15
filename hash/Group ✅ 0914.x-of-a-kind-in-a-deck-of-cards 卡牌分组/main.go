package main

// https://leetcode-cn.com/problems/x-of-a-kind-in-a-deck-of-cards

// ❓存在 int(x) 把数组分为:
// ⚠️ 每组cnt为 int(x)
// ⚠️ 每组val相同

func hasGroupsSizeX(deck []int) bool {
	// 统计元素
	numMpCnt := map[int]int{}
	for _, num := range deck {
		numMpCnt[num]++
	}

	// 最大公约数
	x := -1
	for num := range numMpCnt {
		if x == -1 {
			x = numMpCnt[num]
		} else {
			x = gcd(x, numMpCnt[num])
		}
	}
	return 2 <= x
}

func gcd(x, y int) int {
	if x == 0 {
		return y
	}
	return gcd(y%x, x)
}
