package main

// https://leetcode-cn.com/problems/friends-of-appropriate-ages

// ❓ 适龄的朋友

// age[y] <= 0.5 * age[x] + 7 // 不会向太小的发起请求 >= 15
// age[x] < age[y] 不会向老的发起请求
// 100 < age[y] && age[x] < 100 不会向老的发起请求

func numFriendRequests(ages []int) int {
	var ageMpCnt [121]int32
	for _, age := range ages {
		ageMpCnt[age]++
	}

	var cntReq int32
	var ageMpCntPre [121]int32
	for age := 1; age < 121; age++ {
		cntCur := ageMpCnt[age]
		ageMpCntPre[age] = ageMpCntPre[age-1] + cntCur
		if cntCur == 0 {
			continue
		}

		ageMin := age/2 + 7
		if ageMin < age {
			cntReq += (ageMpCntPre[age] - ageMpCntPre[ageMin] - 1) * cntCur
		}
	}
	return int(cntReq)
}
