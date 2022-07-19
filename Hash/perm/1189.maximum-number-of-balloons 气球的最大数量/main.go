package main

// https://leetcode.cn/problems/maximum-number-of-balloons

// ❓ 气球的最大数量

func maxNumberOfBalloons(text string) int {
	chMpCnt := map[byte]int{}
	for i := range text {
		chMpCnt[text[i]]++
	}

	chMpCntBalloon := map[byte]int{
		'a': 1,
		'b': 1,
		'n': 1,
		'o': 2,
		'l': 2,
	}

	var cntMax = 1<<63 - 1
	for ch := range chMpCntBalloon {
		cnt := chMpCnt[ch] / chMpCntBalloon[ch]
		if cnt < cntMax {
			cntMax = cnt
		}
	}
	return cntMax
}
