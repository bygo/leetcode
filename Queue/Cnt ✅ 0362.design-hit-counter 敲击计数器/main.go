package main

// https://leetcode.cn/problems/design-hit-counter

// ❓ 敲击计数器

type HitCounter struct {
	counts []Counter
}

type Counter struct {
	cnt       int
	timestamp int
}

func Constructor() HitCounter {
	return HitCounter{counts: []Counter{}}
}

// 敲击

func (hc *HitCounter) Hit(timestamp int) {
	cntTop := len(hc.counts) - 1
	if -1 < cntTop && hc.counts[cntTop].timestamp == timestamp {
		// 同一个时间,stack top
		hc.counts[cntTop].cnt++
	} else {
		// 重新计算
		hc.counts = append(hc.counts, Counter{cnt: 1, timestamp: timestamp})
	}
}

// 过去5分钟敲击数

func (hc *HitCounter) GetHits(timestamp int) int {
	var cnt int
	var cntTop = len(hc.counts) - 1
	var start = timestamp - 299
	for {
		// 倒序
		if cntTop < 0 || hc.counts[cntTop].timestamp < start {
			break
		}
		cnt += hc.counts[cntTop].cnt
		cntTop--
	}
	// 优化
	if 0 <= cntTop {
		hc.counts = hc.counts[cntTop:]
	}
	return cnt
}
