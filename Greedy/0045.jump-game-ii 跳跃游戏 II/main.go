package main

// https://leetcode-cn.com/problems/jump-game-ii

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

func maxmiumScore(cards []int, cnt int) int {
	m := [10001]int{}
	for i := range cards {
		m[cards[i]]++
	}

	var odd int
	var oddCnt int
	var res int
	var j int
	var last int
	var lastlast int
	for i := 10000; 0 <= i && j < cnt; i-- {
		if m[i] == 0 {
			continue
		}
		if i%2 == 0 {
			lastlast = last
			last = i * m[i]
			res += last
			j++
		} else {
			if oddCnt == 1 {
				if cnt == j+1 {
					break
				}
				lastCur := last + lastlast
				cur := odd + i*m[i]
				if lastCur < cur {
					res += cur - lastCur
				}
				oddCnt = 0
				if 2 < j {

				}
			} else {
				odd = i * m[i]
				oddCnt = 1
			}
		}
	}
	return res
}
