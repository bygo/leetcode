package main

// https://leetcode-cn.com/problems/shortest-word-distance-ii

type WordDistance struct {
	indexes map[string][]int
}

func Constructor(wordsDict []string) WordDistance {
	indexes := map[string][]int{}
	for i := range wordsDict {
		indexes[wordsDict[i]] = append(indexes[wordsDict[i]], i)
	}
	return WordDistance{
		indexes: indexes,
	}
}

func (wd *WordDistance) Shortest(word1 string, word2 string) int {
	arr1 := wd.indexes[word1]
	arr2 := wd.indexes[word2]
	i, j := len(arr1)-1, len(arr2)-1
	var min = 1<<63 - 1
	for -1 < i && -1 < j {
		var cur int
		if arr1[i] < arr2[j] {
			cur = arr2[j] - arr1[i]
			j--
		} else {
			cur = arr1[i] - arr2[j]
			i--
		}
		if cur < min {
			min = cur
		}
	}
	return min
}
