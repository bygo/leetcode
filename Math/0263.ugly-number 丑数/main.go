package main

// https://leetcode-cn.com/problems/ugly-number

var factors = []int{2, 3, 5}

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	for i := range factors {
		for n%factors[i] == 0 {
			n /= factors[i]
		}
	}

	return n == 1
}
