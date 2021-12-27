package main

// https://leetcode-cn.com/problems/friends-of-appropriate-ages

// ❓ 适龄的朋友
// age[y] <= 0.5 * age[x] + 7 // 不会向太小的发起请求
// age[x] < age[y] 不会向老的发起请求
// 100 < age[y] && age[x] < 100 不会向老的发起请求

func numFriendRequests(ages []int) int {
	var ageMpCnt [121]int
	for _, age := range ages {
		ageMpCnt[age]++
	}

	var cntReq int
	var preAgeMpCnt [121]int
	for age := 1; age < 121; age++ {
		cntCur := ageMpCnt[age]
		preAgeMpCnt[age] = preAgeMpCnt[age-1] + cntCur
		if cntCur == 0 {
			continue
		}

		ageMin := age/2 + 7
		if ageMin < age {
			cntReq += (preAgeMpCnt[age-1] - preAgeMpCnt[ageMin]) * cntCur
			cntReq += cntCur * (cntCur - 1)
		}
	}
	return cntReq
}
