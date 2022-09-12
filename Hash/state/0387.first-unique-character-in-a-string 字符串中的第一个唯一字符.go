package main

// https://leetcode.cn/problems/first-unique-character-in-a-string

// ❓ 字符串中第一个唯一字符索引

func firstUniqCharBit(s string) int {
	sL := len(s)
	chMpIdx := [26]int{}
	// 初始状态
	for i := range chMpIdx {
		chMpIdx[i] = sL
	}

	// 排除
	for i := 0; i < sL; i++ {
		ch := s[i] - 'a'
		if chMpIdx[ch] == sL {
			// 唯一
			chMpIdx[ch] = i
		} else {
			// 重复
			chMpIdx[ch] = sL + 1
		}
	}

	// 挑选
	var idxMinUnique = sL
	for ch := range chMpIdx {
		// 最小唯一索引
		if chMpIdx[ch] < idxMinUnique {
			idxMinUnique = chMpIdx[ch]
		}
	}
	if idxMinUnique < sL {
		return idxMinUnique
	}
	return -1
}

func firstUniqCharQueue(s string) int {
	type pair struct {
		ch  byte
		pos int
	}
	sL := len(s)
	chMpIdx := [26]int{}
	for ch := range chMpIdx {
		chMpIdx[ch] = sL
	}

	var queue []pair
	for i := range s {
		ch := s[i] - 'a'
		if chMpIdx[ch] == sL {
			chMpIdx[ch] = i
			queue = append(queue, pair{
				ch:  ch,
				pos: i,
			})
		} else {
			chMpIdx[ch] = sL + 1
			//for 0 < len(queue) && chMpIdx[queue[0].ch] == sL+1 { // 队列优化map
			//	queue = queue[1:]
			//}
		}
	}
	for 0 < len(queue) && chMpIdx[queue[0].ch] == sL+1 { // 队列优化map
		queue = queue[1:]
	}

	if 0 < len(queue) {
		return queue[0].pos
	}

	return -1
}
