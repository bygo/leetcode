package main

// https://leetcode-cn.com/problems/minimum-index-sum-of-two-lists

// ❓ 两数组相同值的最小索引和 集合
// ⚠️ 没有重复

func findRestaurant(list1 []string, list2 []string) []string {
	strMpIdx := map[string]int{}
	for idx, str := range list1 {
		strMpIdx[str] = idx
	}

	var stringDist []string
	var distMin = 1<<63 - 1
	for idx, str := range list2 {
		_, ok := strMpIdx[str]
		if ok {
			// 存在时计算
			dist := strMpIdx[str] + idx

			if dist < distMin {
				// 比当前还小
				stringDist = []string{str}
				distMin = dist
			} else if dist == distMin {
				// 累加
				stringDist = append(stringDist, str)
			}
		}
	}
	return stringDist
}
