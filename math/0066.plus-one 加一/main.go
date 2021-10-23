package main

func plusOne(digits []int) []int {
	var i = len(digits) - 1
	for 0 <= i && digits[i] == 9 {
		digits[i] = 0
		i--
	}
	if i == -1 {
		return append([]int{1}, digits...)
	}
	digits[i] += 1
	return digits
}
