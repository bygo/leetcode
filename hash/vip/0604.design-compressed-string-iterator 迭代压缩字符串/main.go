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
	compressedString += "#"
	l1 := len(compressedString)
	cnt := 0
	var cur byte
	for i := 0; i < l1; i++ {
		if '0' <= compressedString[i] && compressedString[i] <= '9' {
			cnt = cnt*10 + int(compressedString[i]-'0')
		} else {
			if 0 < cnt {
				q = append(q, pair{
					ch:  cur,
					cnt: cnt,
				})
			}
			cur = compressedString[i]
			cnt = 0
		}
	}

	return StringIterator{
		q: q,
	}
}

func (this *StringIterator) Next() (res byte) {
	if !this.HasNext() {
		return ' '
	}
	this.q[0].cnt--
	res = this.q[0].ch
	if this.q[0].cnt == 0 {
		this.q = this.q[1:]
	}
	return
}

func (this *StringIterator) HasNext() bool {
	return 0 < len(this.q)
}

/**
 * Your StringIterator object will be instantiated and called as such:
 * obj := Constructor(compressedString);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
