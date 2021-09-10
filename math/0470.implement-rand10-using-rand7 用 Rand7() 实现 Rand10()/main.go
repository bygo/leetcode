package main

import "math/rand"

// https://leetcode-cn.com/problems/implement-rand10-using-rand7

func rand10() int {
	var sum int
	for {
		sum = (rand7()-1)*7 + rand7()
		if sum <= 40 {
			return sum%10 + 1
		}

		sum -= 40
		sum = (sum-1)*7 + rand7()
		if sum <= 60 {
			return sum%10 + 1
		}

		sum -= 60
		sum = (sum-1)*7 + rand7()
		if sum <= 20 {
			return sum%10 + 1
		}
	}
}

func rand7() int {
	return rand.Intn(7)
}
