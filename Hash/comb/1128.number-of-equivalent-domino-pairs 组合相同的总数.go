package main

// https://leetcode.cn/problems/number-of-equivalent-domino-pairs

// ❓ 组合相同的数量
// ⚠️ 一个组只有2个数字

func numEquivDominoPairs(d [][]int) int {
	combMpCnt := map[int]int{}
	var cntEquiv int
	for i := range d {
		if d[i][0] < d[i][1] {
			d[i][0], d[i][1] = d[i][1], d[i][0]
		}
		comb := d[i][0]*10 + d[i][1]
		//cnt += combMpCnt[comb]
		combMpCnt[comb]++
	}

	for _, cnt := range combMpCnt {
		cntEquiv += cnt * (cnt - 1) / 2
	}
	return cntEquiv
}
