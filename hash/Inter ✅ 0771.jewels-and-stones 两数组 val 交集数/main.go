package main

// https://leetcode-cn.com/problems/jewels-and-stones/

// ❓ 两数组 val 交集数
func numJewelsInStones(j string, s string) int {
	var chMpCnt = map[byte]int{}
	for i := range j {
		chMpCnt[j[i]] = 1
	}

	var res int
	for i := range s {
		res += chMpCnt[s[i]]
	}
	return res
}
