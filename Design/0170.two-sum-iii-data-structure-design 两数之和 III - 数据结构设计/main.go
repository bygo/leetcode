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

func (this *TwoSum) Add(number int) {
	this.nums[number]++
}

func (this *TwoSum) Find(value int) bool {
	for num, fre := range this.nums {
		if 0 < this.nums[value-num] {
			if value != 2*num || fre == 2 {
				return true
			}
		}
	}
	return false
}

/**
 * Your TwoSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * param_2 := obj.Find(value);
 */
