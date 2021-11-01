package main

// https://leetcode-cn.com/problems/design-compressed-string-iterator

type StringIterator struct {
	q []pair
}

type pair struct {
	ch  byte
	cnt int
}

func Constructor(compressedString string) StringIterator {
	q := []pair{}
	l1 := len(compressedString)
	cnt := 0
	var cur byte
	var i = 0

	f := func() {
		if 0 < cnt {
			q = append(q, pair{
				ch:  cur,
				cnt: cnt,
			})
		}
	}
	for i < l1 {
		if '0' <= compressedString[i] && compressedString[i] <= '9' {
			cnt = cnt*10 + int(compressedString[i]-'0')
		} else {
			f()
			cur = compressedString[i]
			cnt = 0
		}
		i++
	}
	f()
	return StringIterator{
		q: q,
	}
}

func (sI *StringIterator) Next() (res byte) {
	if !sI.HasNext() {
		return ' '
	}
	sI.q[0].cnt--
	res = sI.q[0].ch
	if sI.q[0].cnt == 0 {
		sI.q = sI.q[1:]
	}
	return
}

func (sI *StringIterator) HasNext() bool {
	return 0 < len(sI.q)
}

/**
 * Your StringIterator object will be instantiated and called as such:
 * obj := Constructor(compressedString);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
