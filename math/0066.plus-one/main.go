package main

func plusOne(digits []int) []int {
	l := len(digits)
	carry := 1

	for 0 < l {
		l--
		digits[l] += carry
		if 9 < digits[l] {
			carry = digits[l] / 10
			digits[l] %= 10
		} else {
			carry = 0
		}
	}
	if 0 < carry {
		digits = append(digits, 0)
		copy(digits[1:], digits)
		digits[0] = carry
	}

	return digits
}
