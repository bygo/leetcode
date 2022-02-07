package main

// https://leetcode-cn.com/problems/add-binary

// ❓  二进制字符串求和

func addBinary(a string, b string) string {
	aT := len(a) - 1
	bT := len(b) - 1
	var bufL int
	if aT < bT {
		bufL = bT
	} else {
		bufL = aT
	}
	bufL += 2
	var buf = make([]byte, bufL)
	var bufT = bufL - 1

	var carry int
	for 0 <= aT || 0 <= bT {
		if 0 <= aT {
			carry += int(a[aT] - '0')
			aT--
		}

		if 0 <= bT {
			carry += int(b[bT] - '0')
			bT--
		}

		buf[bufT] = byte(carry%2) + '0'
		carry /= 2
		bufT--
	}

	if 0 < carry {
		buf[bufT] = byte(carry) + '0'
		bufT--
	}
	return string(buf[bufT+1:])
}
