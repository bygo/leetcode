package main

import (
	"container/heap"
	"sort"
)

// https://leetcode-cn.com/problems/course-schedule-iii

func scheduleCourse(courses [][]int) int {
	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})

	h := &Heap{}
	cnt := 0
	for _, course := range courses {
		d := course[0]
		if cnt+d <= course[1] {
			cnt += d
			heap.Push(h, d)
		} else if 0 < h.Len() && d < h.IntSlice[0] {
			cnt += d - h.IntSlice[0]
			h.IntSlice[0] = d
			heap.Fix(h, 0)
		}
	}
	return h.Len()
}

type Heap struct {
	sort.IntSlice
}

func (h Heap) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
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
