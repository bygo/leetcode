package main

// https://leetcode-cn.com/problems/distribute-candies

// ❓ 得到最多类型的糖果
func distributeCandies(candy []int) int {
	candyL := len(candy)
	typMp := map[int]struct{}{}
	for _, typ := range candy {
		typMp[typ] = struct{}{} // 类型数
	}

	typMpL := len(typMp)
	if candyL/2 < typMpL { // 类型数超过 一半，每样一颗
		return candyL / 2
	}
	return typMpL // 小于一半 最多typMpL
}
