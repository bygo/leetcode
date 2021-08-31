package main

// https://leetcode-cn.com/problems/two-sum-iii-data-structure-design

type TwoSum struct {
	nums map[int]int
}

/** Initialize your data structure here. */
func Constructor() TwoSum {
	return TwoSum{
		nums: map[int]int{},
	}
}

/** Add the number to an internal data structure.. */
func (this *TwoSum) Add(number int) {
	this.nums[number]++
}

/** Find if there exists any pair of numbers which sum is equal to the value. */
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
