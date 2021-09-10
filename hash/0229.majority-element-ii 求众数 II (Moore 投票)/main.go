package main

// https://leetcode-cn.com/problems/majority-element-ii

type candidate struct {
	num int
	cnt int
}

func majorityElement(nums []int) []int {
	const base = 3
	var cs [base - 1]candidate
	for i := range nums {
		for j := range cs {
			if cs[j].num == nums[i] {
				cs[j].cnt++
				goto END
			}
		}
		for j := range cs {
			if cs[j].cnt == 0 {
				cs[j].num = nums[i]
				cs[j].cnt = 1
				goto END
			}
		}

		for j := range cs {
			cs[j].cnt--
		}
	END:
	}

	for i := range cs {
		cs[i].cnt = 0
	}
	for i := range nums {
		for j := range cs {
			if cs[j].num == nums[i] {
				cs[j].cnt++
				break
			}
		}
	}

	var res []int
	f := len(nums) / base
	for i := range cs {
		if f < cs[i].cnt {
			res = append(res, cs[i].num)
		}
	}
	return res
}
