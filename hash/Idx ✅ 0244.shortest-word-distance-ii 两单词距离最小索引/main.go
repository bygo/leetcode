package main

// https://leetcode-cn.com/problems/shortest-word-distance-ii

// ❓ 两单词距离最小索引

type WordDistance struct {
	strMpIdxes map[string][]int
}

func Constructor(wordsDict []string) WordDistance {
	strMpIdxes := map[string][]int{}
	for i, str := range wordsDict {
		strMpIdxes[str] = append(strMpIdxes[str], i)
	}
	return WordDistance{
		strMpIdxes: strMpIdxes,
	}
}

func (wd *WordDistance) Shortest(word1 string, word2 string) int {
	idxes1 := wd.strMpIdxes[word1]
	idxes2 := wd.strMpIdxes[word2]
	i, j := len(idxes1)-1, len(idxes2)-1
	var distMin = 1<<63 - 1
	var distTmp int
	for -1 < i && -1 < j {
		if idxes1[i] < idxes2[j] {
			// 2逼近1
			distTmp = idxes2[j] - idxes1[i]
			j--
		} else {
			// 1逼近2
			distTmp = idxes1[i] - idxes2[j]
			i--
		}

		// 距离
		if distTmp < distMin {
			distMin = distTmp
		}
	}
	return distMin
}
