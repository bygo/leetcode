package main

func subtractProductAndSum(n int) int {
	x, y := 1, 0
	for 0 < n {
		carry := n % 10
		n /= 10
		x *= carry
		y += carry
	}
	return x - y
}
