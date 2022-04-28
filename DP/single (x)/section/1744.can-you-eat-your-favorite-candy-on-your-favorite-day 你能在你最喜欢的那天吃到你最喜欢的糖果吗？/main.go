package main

// https://leetcode-cn.com/problems/can-you-eat-your-favorite-candy-on-your-favorite-day

func canEat(candiesCount []int, queries [][]int) []bool {
	var bools []bool
	cL := len(candiesCount)
	for typ := 1; typ < cL; typ++ {
		candiesCount[typ] += candiesCount[typ-1]
	}
	var lo, hi int

	for _, q := range queries {
		typ := q[0]
		day := q[1]
		cntCap := q[2]

		hiEat := day + 1        // 每天一颗
		loEat := cntCap * hiEat // 每天最大

		hi = candiesCount[typ]
		if 0 == typ {
			lo = 1
		} else {
			lo = candiesCount[typ-1] + 1
		}

		// 最大够不到最左  最小超过最右 都为false
		//bs = append(bs, !(max < lo || hi < min))
		bools = append(bools, lo <= loEat && hiEat <= hi)
	}
	return bools
}
