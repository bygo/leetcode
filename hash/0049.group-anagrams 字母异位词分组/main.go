package main

// https://leetcode-cn.com/problems/group-anagrams

func groupAnagrams(strs []string) [][]string {
	m := map[[26]int][]string{}
	for i := range strs {
		cnt := [26]int{}
		for j := range strs[i] {
			cnt[strs[i][j]-'a']++
		}
		m[cnt] = append(m[cnt], strs[i])
	}
	res := [][]string{}
	for i := range m {
		res = append(res, m[i])
	}
	return res
}
