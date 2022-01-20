package main

// https://leetcode-cn.com/problems/add-binary

// ❓  二进制字符串求和

func addBinary(a string, b string) string {
	var carry int
	var buf []byte
	aTop := len(a) - 1
	bTop := len(b) - 1

	for 0 <= aTop || 0 <= bTop || 0 < carry {
		var num = carry
		if 0 <= aTop {
			if a[aTop] == '1' {
				num++
			}
			aTop--
		}
		if 0 <= bTop {
			if b[bTop] == '1' {
				num++
			}
			bTop--
		}
		if num%2 == 0 {
			buf = append(buf, '0')
		} else {
			buf = append(buf, '1')
		}
		carry = num / 2
	}

	lo, hi := 0, len(buf)-1
	for lo < hi {
		buf[lo], buf[hi] = buf[hi], buf[lo]
		lo++
		hi--
	}
	return string(buf)
}
