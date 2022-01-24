package main

// https://leetcode-cn.com/problems/add-binary

// ❓  二进制字符串求和

func addBinary(a string, b string) string {
	var carry int
	aTop := len(a) - 1
	bTop := len(b) - 1
	var bufL int
	if aTop < bTop {
		bufL = bTop + 2
	} else {
		bufL = aTop + 2
	}
	var buf = make([]byte, bufL)
	var bufTop = bufL - 1
	for 0 <= aTop || 0 <= bTop {
		if 0 <= aTop {
			carry += int(a[aTop] - '0')
			aTop--
		}

		if 0 <= bTop {
			carry += int(b[bTop] - '0')
			bTop--
		}

		buf[bufTop] = byte(carry%2) + '0'
		carry /= 2
		bufTop--
	}

	if 0 < carry {
		buf[bufTop] = byte(carry) + '0'
		bufTop--
	}
	return string(buf[bufTop+1:])
}
