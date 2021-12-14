package main

// https://leetcode-cn.com/problems/jewels-and-stones/

func numJewelsInStones(j string, s string) int {
	var m = map[byte]int{}
	for i := range j {
		m[j[i]] = 1
	}

	var res int
	for i := range s {
		res += m[s[i]]
	}
	return res
}
