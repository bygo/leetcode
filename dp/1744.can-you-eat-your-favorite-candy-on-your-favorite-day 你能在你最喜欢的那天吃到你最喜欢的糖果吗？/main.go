package main

// Link: https://leetcode-cn.com/problems/can-you-eat-your-favorite-candy-on-your-favorite-day

// 前缀和
func canEat(candiesCount []int, queries [][]int) []bool {
	var res []bool

	for i := 1; i < len(candiesCount); i++ {
		candiesCount[i] += candiesCount[i-1]
	}

	for _, q := range queries {
		min := q[1] + 1
		max := q[2] * min

		var left int
		var right = candiesCount[q[0]]
		if 0 == q[0] {
			left = 1
		} else {
			left = candiesCount[q[0]-1] + 1
		}

		// 最大够不到最左  最小超过最右 都为false
		//res = append(res, !(max < left || right < min))
		res = append(res, left <= max && min <= right)
	}
	return res
}
