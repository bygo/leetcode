package main

// https://leetcode.cn/problems/friends-of-appropriate-ages

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

func shiftingLetters(s string, shifts [][]int) string {
	sL := len(s)
	diff := make([]int, sL+1)
	for _, shift := range shifts {
		shift[2] = shift[2]<<1 - 1
		diff[shift[0]] += shift[2]
		diff[shift[1]+1] -= shift[2]
	}

	buf := []byte(s)
	var offset int
	for i := range buf {
		offset = (diff[i]+offset)%26 + 26
		buf[i] = (buf[i]-'a'+byte(offset))%26 + 'a'
	}
	return string(buf)
}
