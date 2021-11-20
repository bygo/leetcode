package main

// https://leetcode-cn.com/problems/majority-element-ii

type candidate struct {
	num, cnt int
}

const base = 3
const csLen = base - 1

func majorityElement(nums []int) []int {
	var l = len(nums)
	var cs [csLen]candidate
	for _, num := range nums {
		var firstSlot = -1
		var k int
		for k < csLen && cs[k].num != num {
			if cs[k].cnt == 0 {
				firstSlot = k
			}
			k++
		}
		if k < csLen {
			cs[k].cnt++
		} else if k == csLen {
			if -1 < firstSlot {
				cs[firstSlot].num = num
				cs[firstSlot].cnt = 1
			} else {
				for i := range cs {
					cs[i].cnt--
				}
			}
		}
	}
	for i := range cs {
		cs[i].cnt = 0
	}
	for _, num := range nums {
		for i := range cs {
			if cs[i].num == num {
				cs[i].cnt++
				break
			}
		}
	}

	var limit = l / base
	var res []int
	for i := range cs {
		if limit < cs[i].cnt {
			res = append(res, cs[i].num)
		}
	}
	return res
}
