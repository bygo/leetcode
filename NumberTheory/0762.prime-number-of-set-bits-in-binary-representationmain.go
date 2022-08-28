package main

import "math/bits"

// https://leetcode.cn/problems/prime-number-of-set-bits-in-binary-representation

func isPrime(x int) bool {
	if x < 2 {
		return false
	}
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func countPrimeSetBits(left, right int) int {
	var cnt int
	for x := left; x <= right; x++ {
		if isPrime(bits.OnesCount(uint(x))) {
			cnt++
		}
	}
	return cnt
}
