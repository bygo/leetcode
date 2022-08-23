package add

// https://leetcode.cn/problems/add-binary

// ❓  二进制字符串求和

func addBinary(a string, b string) string {
	aT := len(a) - 1
	bT := len(b) - 1
	if bT < aT {
		aT, bT = bT, aT
		a, b = b, a
	}
	lo := bT + 1
	var bufL = lo + 1
	var buf = make([]byte, bufL)
	var carry byte
	for 0 <= aT {
		carry += a[aT] + b[bT] - '0'*2
		aT--
		bT--
		buf[lo] = carry%2 + '0'
		carry /= 2
		lo--
	}

	for 0 <= bT {
		carry += b[bT] - '0'
		bT--
		buf[lo] = carry%2 + '0'
		carry /= 2
		lo--
	}

	if 0 < carry {
		buf[lo] = carry + '0'
		lo--
	}
	return string(buf[lo+1:])
}
