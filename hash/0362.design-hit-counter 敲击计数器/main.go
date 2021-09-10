package main

// https://leetcode-cn.com/problems/design-hit-counter

type HitCounter struct {
	m []Counter
}

type Counter struct {
	cnt       int
	timestamp int
}

func Constructor() HitCounter {
	return HitCounter{m: []Counter{}}
}

func (this *HitCounter) Hit(timestamp int) {
	n := len(this.m) - 1
	if -1 < n && this.m[n].timestamp == timestamp {
		this.m[n].cnt++
	} else {
		this.m = append(this.m, Counter{cnt: 1, timestamp: timestamp})
	}
}

func (this *HitCounter) GetHits(timestamp int) int {
	var res int
	var t int
	var start = timestamp - 299
	for i := range this.m {
		if this.m[i].timestamp < start {
			t++
		} else {
			res += this.m[i].cnt
		}
	}
	this.m = this.m[t:]
	return res
}
