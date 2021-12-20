package main

func mul(a, b int) int {
	var res = 0
	var cnt int
	if a < 0 {
		cnt++
		a = -a
	}
	if b < 0 {
		cnt++
		b = -b
	}
	for 0 < b {
		if 1 == b&1 {
			res += a
		}
		b >>= 1
		a <<= 1
	}
	if cnt%2 == 1 {
		return -res
	}
	return res
}
