package main

// https://leetcode-cn.com/problems/degree-of-an-array

type E struct {
	cnt int
	l   int
	r   int
}

func findShortestSubArray(nums []int) int {
	m := map[int]*E{}
	for i := range nums {
		_, ok := m[nums[i]]
		if ok {
			m[nums[i]].r = i
			m[nums[i]].cnt++
		} else {
			m[nums[i]] = &E{
				cnt: 1,
				l:   i,
				r:   i,
			}
		}
	}

	var cnt int
	var res int
	for _, e := range m {
		if cnt < e.cnt {
			cnt = e.cnt
			res = e.r - e.l + 1
		} else if cnt == e.cnt {
			cur := e.r - e.l + 1
			if cur < res {
				res = cur
			}
		}
	}
	return res
}
