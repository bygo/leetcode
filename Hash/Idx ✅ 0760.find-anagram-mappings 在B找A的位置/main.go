package main

// https://leetcode-cn.com/problems/find-anagram-mappings

// ❓ 找出 A里面元素，在B里的索引值

func anagramMappings(A []int, B []int) []int {
	// B 的位置
	var numMpIdx = map[int]int{}
	for idx, num := range B {
		numMpIdx[num] = idx
	}

	// 索引值
	var idxesA []int
	for _, num := range A {
		idxesA = append(idxesA, numMpIdx[num])
	}
	return idxesA
}
