package main

func findMagicIndex(nums []int) int {
	var dfs func(l, r int) int
	dfs = func(l, r int) int {
		if r < l {
			return -1
		}
		mid := l + (r-l)/2
		res := dfs(l, mid-1)
		if res != -1 {
			return res
		} else if nums[mid] == mid {
			return mid
		}
		return dfs(mid+1, r)
	}
	return dfs(0, len(nums)-1)
}

func findMagicIndex(nums []int) int {
	i := 0
	for i < len(nums) {
		if i == nums[i] {
			return i
		} else if i < nums[i] {
			i = nums[i]
		} else {
			i++
		}
	}
	return -1
}
