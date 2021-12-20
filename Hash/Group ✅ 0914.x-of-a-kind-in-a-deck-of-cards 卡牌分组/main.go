package main

// https://leetcode-cn.com/problems/x-of-a-kind-in-a-deck-of-cards

// ❓存在 int(x) 把数组分为:
// ⚠️ 每组cnt为 int(g)
// ⚠️ 每组val相同

func hasGroupsSizeX(deck []int) bool {
	// 统计元素
	numMpCnt := map[int]int{}
	for _, num := range deck {
		numMpCnt[num]++
	}

	// 最大公约数
	g := -1
	for num := range numMpCnt {
		if g == -1 {
			g = numMpCnt[num]
		} else {
			g = gcd(g, numMpCnt[num])
		}
	}
	return 2 <= g
}

func gcd(x, y int) int {
	if x == 0 {
		return y
	}
	return gcd(y%x, x)
}
