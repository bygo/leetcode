package main

// https://leetcode-cn.com/problems/add-binary

// ❓  二进制字符串求和

func addBinary(a string, b string) string {
	aT := len(a) - 1
	bT := len(b) - 1
	if bT < aT {
		aT, bT = bT, aT
		a, b = b, a
	}
	bufT := bT + 1
	var bufL = bufT + 1
	var buf = make([]byte, bufL)
	var carry byte
	for 0 <= aT {
		carry += a[aT] + b[bT] - '0' - '0'
		aT--
		bT--
		buf[bufT] = carry%2 + '0'
		carry /= 2
		bufT--
	}

	for 0 <= bT {
		carry += b[bT] - '0'
		bT--
		buf[bufT] = carry%2 + '0'
		carry /= 2
		bufT--
	}

	if 0 < carry {
		buf[bufT] = carry + '0'
		bufT--
	}
	return string(buf[bufT+1:])
}
