package main

// https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted

func twoSum(numbers []int, target int) []int {
	l1 := len(numbers)
	for i := 0; i < l1; i++ {
		l, r := i+1, l1-1
		for l <= r {
			mid := (r-l)/2 + l
			cur := target - numbers[i]
			if numbers[mid] == cur {
				return []int{i + 1, mid + 1}
			} else if cur < numbers[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return []int{-1, -1}
}

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l < r {
		sum := numbers[l] + numbers[r]
		if sum == target {
			return []int{l + 1, r + 1}
		} else if sum < target {
			l++
		} else if target < sum {
			r--
		}
	}
	return []int{-1, -1}
}
