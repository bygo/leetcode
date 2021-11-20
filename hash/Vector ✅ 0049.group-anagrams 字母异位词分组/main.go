package main

// https://leetcode-cn.com/problems/group-anagrams

func groupAnagrams(strs []string) [][]string {
	m := map[[26]int][]string{}
	for i := range strs {
		var vector [26]int
		for j := range strs[i] {
			vector[strs[i][j]-'a']++
		}
		m[vector] = append(m[vector], strs[i])
	}
	var res [][]string
	for i := range m {
		res = append(res, m[i])
	}
	return res
}
