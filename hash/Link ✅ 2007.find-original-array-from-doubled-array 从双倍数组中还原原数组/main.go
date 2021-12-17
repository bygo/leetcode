package main

// https://leetcode-cn.com/problems/find-original-array-from-doubled-array

func findOriginalArray(changed []int) []int {
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
