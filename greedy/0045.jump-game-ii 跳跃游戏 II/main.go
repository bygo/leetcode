package main

// fewest steps jumps to point ?
// https://leetcode-cn.com/problems/jump-game-ii

// find farthest point where have must jump
// init 0
func jump(nums []int) int {
	var steps, farthest, point int
	for i, num := range nums {
		point = max(point, i+num)
		if i == farthest {
			farthest = point
			steps++
		}
	}
	return steps
}

// find first step jump point
func jump(nums []int) int {
	var steps int
	point := len(nums) - 1
	for 0 < point {
		for i := 0; i < point; i++ {
			if point <= i+nums[i] {
				point = i
				steps++
				break
			}
		}
	}
	return steps
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
