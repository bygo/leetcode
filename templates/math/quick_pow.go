package main

func pow(x float64, n int) float64 {
	if x == 1.0 || n == 0 {
		return 1.0
	}
	var res = 1.0
	var y = n
	if n < 0 {
		y = -y
	}
	for 0 < y {
		if y&1 == 1 {
			res *= x
		}
		x *= x
		y >>= 1
	}
	if n < 0 {
		return 1.0 / res
	}
	return res
}
