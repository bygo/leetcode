package main

import "math/rand"

// https://leetcode-cn.com/problems/shuffle-an-array

type Solution struct {
	ori, cur []int
}

func Constructor(nums []int) Solution {
	return Solution{
		ori: nums,
		cur: append([]int{}, nums...),
	}
}

func (s *Solution) Reset() []int {
	return s.ori
}

func (s *Solution) Shuffle() []int {
	n := len(s.cur)
	for i := range s.cur {
		j := i + rand.Intn(n-i)
		s.cur[i], s.cur[j] = s.cur[j], s.cur[i]
	}
	return s.cur
}

type Solution struct {
	ori, cur []int
}

func Constructor(nums []int) Solution {
	return Solution{
		ori: nums,
		cur: append([]int{}, nums...),
	}
}
