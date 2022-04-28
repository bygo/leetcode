package main

// https://leetcode-cn.com/problems/number-of-boomerangs

func numberOfBoomerangs(points [][]int) int {
	var cntRes int
	for _, p := range points {
		distMpCnt := map[int]int{}
		for _, q := range points {
			// p 到 q 的距离
			dist := (p[0]-q[0])*(p[0]-q[0]) + (p[1]-q[1])*(p[1]-q[1])
			cntRes += distMpCnt[dist] * 2
			distMpCnt[dist]++
		}

		//for _, cnt := range distMpCnt {
		//	cntRes += cnt * (cnt - 1)
		//}
	}
	return cntRes
}
