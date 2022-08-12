package main

// https://leetcode.cn/problems/design-compressed-string-iterator

// ❓ 迭代压缩字符串
// ⚠️ L1e2t1C1o1d1e1

type StringIterator struct {
	queue []pair
}

type pair struct {
	ch  byte
	cnt int
}

func Constructor(compressedString string) StringIterator {
	var queue []pair
	compressedStringL := len(compressedString)
	cnt := 0
	var chCur byte
	var idx = 0

	fn := func() {
		if 0 < cnt {
			queue = append(queue, pair{
				ch:  chCur,
				cnt: cnt,
			})
		}
	}
	for idx < compressedStringL {
		if '0' <= compressedString[idx] && compressedString[idx] <= '9' {
			cnt = cnt*10 + int(compressedString[idx]-'0')
		} else {
			fn()
			chCur = compressedString[idx]
			cnt = 0
		}
		idx++
	}
	fn()
	return StringIterator{
		queue: queue,
	}
}

func (sI *StringIterator) Next() byte {
	if !sI.HasNext() {
		return ' '
	}
	sI.queue[0].cnt--
	ch := sI.queue[0].ch
	if sI.queue[0].cnt == 0 {
		sI.queue = sI.queue[1:]
	}
	return ch
}

func (sI *StringIterator) HasNext() bool {
	return 0 < len(sI.queue)
}
