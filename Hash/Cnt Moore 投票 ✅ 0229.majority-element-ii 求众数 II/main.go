package main

// https://leetcode-cn.com/problems/majority-element-ii

// ❓ 超过1/n的数

type Elem struct {
	num, cnt int
}

const base = 3
const elemL = base - 1

func majorityElement(nums []int) []int {
	var l = len(nums)
	var elems [elemL]Elem
	for _, num := range nums {
		var slotInsert = -1
		var idx int
		// 找出第一个插槽
		for idx < elemL {
			if elems[idx].num == num {
				// 同值插槽
				break
			} else if elems[idx].cnt == 0 {
				// 第一个空槽
				slotInsert = idx
			}
			idx++
		}

		if idx < elemL {
			// 同值插槽
			elems[idx].cnt++
		} else if idx == elemL {
			if -1 < slotInsert {
				// 有空槽 替换
				elems[slotInsert].num = num
				elems[slotInsert].cnt = 1
			} else {
				// 全部自减
				for i := range elems {
					elems[i].cnt--
				}
			}
		}
	}

	// 置零
	for i := range elems {
		elems[i].cnt = 0
	}

	// 因为抵消，需要重新计数
	for _, num := range nums {
		for i := range elems {
			if elems[i].num == num {
				elems[i].cnt++
				break
			}
		}
	}

	// 多个答案
	var limit = l / base
	var numsMode []int
	for i := range elems {
		if limit < elems[i].cnt {
			numsMode = append(numsMode, elems[i].num)
		}
	}
	return numsMode
}
