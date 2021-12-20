package main

// https://leetcode-cn.com/problems/find-anagram-mappings

// ❓ 在B找A的位置

func anagramMappings(A []int, B []int) []int {
	// B 的位置
	var numMpIdx = map[int]int{}
	for idx, num := range B {
		numMpIdx[num] = idx
	}

	var idxesA []int
	for _, num := range A {
		idxesA = append(idxesA, numMpIdx[num])
	}
	return idxesA
}
