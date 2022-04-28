package main

// https://leetcode-cn.com/problems/two-sum-iii-data-structure-design

type TwoSum struct {
	nums map[int]int
}

func Constructor() TwoSum {
	return TwoSum{
		nums: map[int]int{},
	}
}

func (ts *TwoSum) Add(num int) {
	ts.nums[num]++
}

func (ts *TwoSum) Find(target int) bool {
	for num, cnt := range ts.nums {
		if 0 < ts.nums[target-num] {
			// 0
			if target != 2*num || cnt == 2 {
				return true
			}
		}
	}
	return false
}
