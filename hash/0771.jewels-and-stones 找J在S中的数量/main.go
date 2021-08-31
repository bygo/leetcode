package main

// https://leetcode-cn.com/problems/jewels-and-stones/


func numJewelsInStones(J string, S string) int {
	var m = map[byte]int{}
	for i := range J {
		m[J[i]] = 1
	}

	var res int
	for i := range S {
		res += m[S[i]]
	}
	return res
}
