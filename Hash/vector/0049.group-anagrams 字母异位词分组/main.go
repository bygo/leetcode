package main

// https://leetcode-cn.com/problems/group-anagrams

// ❓ 异位词分组

func groupAnagrams(strs []string) [][]string {
	vectorMpStrings := map[[26]int][]string{}
	for _, str := range strs {
		var vector [26]int
		for j := range str {
			ch := str[j] - 'a'
			vector[ch]++
		}
		vectorMpStrings[vector] = append(vectorMpStrings[vector], str)
	}
	var vectorsStrings [][]string
	for _, strings := range vectorMpStrings {
		vectorsStrings = append(vectorsStrings, strings)
	}
	return vectorsStrings
}
