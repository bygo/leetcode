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

func (hc *HitCounter) Hit(timestamp int) {
	n := len(hc.m) - 1
	if -1 < n && hc.m[n].timestamp == timestamp {
		hc.m[n].cnt++
	} else {
		hc.m = append(hc.m, Counter{cnt: 1, timestamp: timestamp})
	}
}

func (hc *HitCounter) GetHits(timestamp int) int {
	var res int
	var t int
	var start = timestamp - 299
	for i := range hc.m {
		if hc.m[i].timestamp < start {
			t++
		} else {
			res += hc.m[i].cnt
		}
	}
	hc.m = hc.m[t:]
	return res
}
