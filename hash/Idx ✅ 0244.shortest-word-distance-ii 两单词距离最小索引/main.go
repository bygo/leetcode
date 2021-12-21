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

func (wd *WordDistance) Shortest(str1 string, str2 string) int {
	idxes1 := wd.strMpIdxes[str1]
	idxes2 := wd.strMpIdxes[str2]
	idx1, idx2 := len(idxes1)-1, len(idxes2)-1
	var distMin = 1<<63 - 1
	var distTmp int
	for -1 < idx1 && -1 < idx2 {
		if idxes1[idx1] < idxes2[idx2] {
			// 2逼近1
			distTmp = idxes2[idx2] - idxes1[idx1]
			idx2--
		} else {
			// 1逼近2
			distTmp = idxes1[idx1] - idxes2[idx2]
			idx1--
		}

		// 距离
		if distTmp < distMin {
			distMin = distTmp
		}
	}
	return distMin
}
