package main

// https://leetcode-cn.com/problems/shortest-word-distance-ii

type WordDistance struct {
	m map[string][]int
}

func Constructor(wordsDict []string) WordDistance {
	m := map[string][]int{}
	for i := range wordsDict {
		m[wordsDict[i]] = append(m[wordsDict[i]], i)
	}
	return WordDistance{
		m: m,
	}
}

func (this *WordDistance) Shortest(word1 string, word2 string) int {
	arr1 := this.m[word1]
	arr2 := this.m[word2]
	i, j := len(arr1)-1, len(arr2)-1
	var min = 1<<63 - 1
	for 0 <= i && 0 <= j {
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

/**
 * Your WordDistance object will be instantiated and called as such:
 * obj := Constructor(wordsDict);
 * param_1 := obj.Shortest(word1,word2);
 */
