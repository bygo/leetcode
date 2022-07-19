package main

import "sort"

// https://leetcode.cn/problems/friends-of-appropriate-ages

// ❓ 适龄的朋友

// age[y] <= 0.5 * age[x] + 7 // 不会向太小的发起请求 >= 15
// age[x] < age[y] 不会向老的发起请求
// 100 < age[y] && age[x] < 100 不会向老的发起请求

func numFriendRequests(ages []int) int {
	var cntReq int
	sort.Ints(ages)
	idxLeft, idxRight := 0, 0
	agesL := len(ages)
	agesTop := agesL - 1
	for _, age := range ages {
		if age < 15 {
			continue
		}
		// 左边界
		for ages[idxLeft] <= age/2+7 {
			idxLeft++
		}

		// 右边界,如果出现相同 idxRight 会超过自身长度
		for idxRight < agesTop && ages[idxRight+1] <= age {
			idxRight++
		}
		cntReq += idxRight - idxLeft
	}
	return cntReq
}
