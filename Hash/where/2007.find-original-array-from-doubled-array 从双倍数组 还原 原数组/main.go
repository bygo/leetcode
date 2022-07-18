package main

// https://leetcode.cn/problems/find-original-array-from-doubled-array

// ❓ 从 双倍数组（翻倍后加入原数组） 还原 原数组
// ⚠️ 0 ~ 10^5

func findOriginalArray(changed []int) []int {
	// 奇数无法还原
	if len(changed)%2 == 1 {
		return nil
	}
	// 计数
	numMpCnt := [100001]int{}
	var cnt int
	for _, num := range changed {
		numMpCnt[num]++
		cnt++
	}

	var numsOri []int
	for num := 0; num <= 50000; num++ {
		// 必须从小开始抵消
		for 0 < numMpCnt[num] {
			numMpCnt[num]--
			numMpCnt[num*2]--
			cnt -= 2
			numsOri = append(numsOri, num)
			if numMpCnt[num*2] < 0 {
				return nil
			}
		}
	}
	if 0 < cnt {
		return nil
	}
	return numsOri
}
