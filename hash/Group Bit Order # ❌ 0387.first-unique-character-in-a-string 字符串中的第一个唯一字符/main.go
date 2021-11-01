package main

// https://leetcode-cn.com/problems/first-unique-character-in-a-string

func firstUniqCharMark(s string) int {
	m := [26]int{}
	for i := range s {
		m[s[i]-'a']++
	}

	for i := range s {
		if m[s[i]] == 1 {
			return i
		}
	}
	return -1
}

func firstUniqCharBit(s string) int {
	l1 := len(s)
	m := [26]int{}
	for i := range m {
		m[i] = l1
	}
	for i := 0; i < l1; i++ {
		ch := s[i] - 'a'
		if m[ch] == l1 {
			m[ch] = i
		} else {
			m[ch] = l1 + 1
		}
	}

	var res = l1
	for i := range m {
		if m[i] < res {
			res = m[i]
		}
	}
	if res < l1 {
		return res
	}
	return -1
}

func firstUniqCharQueue(s string) int {
	type pair struct {
		ch  byte
		pos int
	}
	l1 := len(s)
	m := [26]int{}
	for i := range m {
		m[i] = l1
	}

	var queue []pair
	for i := range s {
		ch := s[i] - 'a'
		if m[ch] == l1 {
			m[ch] = i
			queue = append(queue, pair{
				ch:  ch,
				pos: i,
			})
		} else {
			m[ch] = l1 + 1
			//for 0 < len(queue) && m[queue[0].ch] == l1+1 { // 队列优化map
			//	queue = queue[1:]
			//}
		}
	}
	for 0 < len(queue) && m[queue[0].ch] == l1+1 { // 队列优化map
		queue = queue[1:]
	}

	if 0 < len(queue) {
		return queue[0].pos
	}

	return -1
}
