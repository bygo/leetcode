package main

// https://leetcode-cn.com/problems/intersection-of-two-arrays

// ❓ 相对交集 val 取唯一值

func intersection(nums1 []int, nums2 []int) []int {
	var numMp = map[int]struct{}{}
	for _, num := range nums1 {
		numMp[num] = struct{}{}
	}

	var numsInter []int
	for _, num := range nums2 {
		_, ok := numMp[num]
		if ok {
			numsInter = append(numsInter, num)
			delete(numMp, num)
		}
	}
	return numsInter
}
