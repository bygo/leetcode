package main

// https://leetcode-cn.com/problems/circular-array-loop

func circularArrayLoop(nums []int) bool {
	n := len(nums)
	next := func(cur int) int {
		// 当前位置 + 步长，步长为负，向上取同余数
		return ((cur+nums[cur])%n + n) % n
	}

	for i, num := range nums {
		if num != 0 {
			for slow, fast := i, next(i);
				0 < nums[slow]*nums[fast] && 0 < nums[slow]*nums[next(fast)];
			slow, fast = next(slow), next(next(fast)) {
				if slow == fast {
					if slow == next(slow) { // 步长 0
						break
					}
					return true
				}
			}
			for l, r := i, next(i); 0 < nums[l]*nums[r]; l, r = r, next(r) { // 走过的忽略
				nums[l] = 0
			}
		}
	}
	return false
}
