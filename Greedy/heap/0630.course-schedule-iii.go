package heap

import (
	"container/heap"
	"sort"
)

// https://leetcode.cn/problems/course-schedule-iii

func scheduleCourse(courses [][]int) int {
	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})

	h := &Heap{}
	total := 0
	for _, course := range courses {
		duration := course[0]
		deadline := course[1]
		if total+duration <= deadline {
			// 贪心时间足够
			total += duration
			heap.Push(h, duration)
		} else if 0 < h.Len() && duration < h.IntSlice[0] {
			// 更短的时间
			total += duration - h.IntSlice[0]
			h.IntSlice[0] = duration
			heap.Fix(h, 0)
		}
	}
	return h.Len()
}

type Heap struct {
	sort.IntSlice
}

func (h Heap) Less(i, j int) bool {
	return h.IntSlice[j] < h.IntSlice[i]
}

func (h *Heap) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *Heap) Pop() interface{} {
	a := h.IntSlice
	x := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return x
}
