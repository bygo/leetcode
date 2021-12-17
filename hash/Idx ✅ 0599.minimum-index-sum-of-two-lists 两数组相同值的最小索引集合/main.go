package main

// https://leetcode-cn.com/problems/minimum-index-sum-of-two-lists

// ❓ 两数组相同值的最小索引集合
// ⚠️ 没有重复

func findRestaurant(list1 []string, list2 []string) []string {
	strMpIdx := map[string]int{}
	for i, str := range list1 {
		strMpIdx[str] = i + 1
	}

	var stringsDist []string
	var distMin = 1<<63 - 1
	for i := range list2 {
		if 0 < strMpIdx[list2[i]] {
			// 存在时计算
			dist := strMpIdx[list2[i]] + i

			if dist < distMin {
				// 比当前还小
				stringsDist = []string{list2[i]}
				distMin = dist
			} else if dist == distMin {
				// 累加
				stringsDist = append(stringsDist, list2[i])
			}
		}
	}
	return stringsDist
}
