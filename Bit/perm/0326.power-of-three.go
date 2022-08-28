package main

func isPowerOfThree(n int) bool {
	for 1 < n && n%3 == 0 {
		n /= 3
	}
	return n == 1
}

func isPowerOfThree(n int) bool {
	return 0 < n && 0b1000101010001101011001111011011%n == 0 // TODO 约数
}
