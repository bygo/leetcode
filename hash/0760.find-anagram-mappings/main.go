package main

//Title: Find Anagram Mappings
// Link: https://leetcode-cn.com/problems/find-anagram-mappings

func anagramMappings(A []int, B []int) []int {
	var m = map[int]int{}
	for k, v := range B {
		m[v] = k
	}

	var res []int
	for _, v := range A {
		res = append(res, m[v])
	}
	return res
}
