package main

// https://leetcode-cn.com/problems/circular-array-loop

func circularArrayLoop(nums []int) bool {
	n := len(nums)
	next := func(cur int) int {
		return ((cur+nums[cur])%n + n) % n
	}

	for i, num := range nums {
		if num != 0 {
			slow, fast := i, next(i)
			for 0 < nums[slow]*nums[fast] && 0 < nums[slow]*nums[next(fast)] {
				if slow == fast {
					if slow == next(slow) {
						break
					}
					return true
				}
				slow = next(slow)
				fast = next(next(fast))
			}
			j := i
			for 0 < nums[j]*nums[next(j)] {
				tmp := j
				j = next(j)
				nums[tmp] = 0
			}
		}

	}
	return false
}
