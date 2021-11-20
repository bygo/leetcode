package main

// https://leetcode-cn.com/problems/combinations

func combine(n int, k int) [][]int {
	var res [][]int
	var cur []int
	for i := 1; i <= k; i++ {
		cur = append(cur, i)
	}
	cur = append(cur, n+1)

	var j = 0
	for j < k {
		tmp := make([]int, k)
		copy(tmp, cur[:k])
		res = append(res, tmp)
		j = 0
		for j < k && cur[j]+1 == cur[j+1] {
			cur[j] = j + 1
			j++
		}
		cur[j] ++
	}
	return res
}
