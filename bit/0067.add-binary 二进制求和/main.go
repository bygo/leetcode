package main

// https://leetcode-cn.com/problems/add-binary

func addBinary(a string, b string) string {
	carry := 0
	var i, j = len(a) - 1, len(b) - 1
	var res = []byte{}
	for 0 <= i || 0 <= j || carry != 0 {
		if 0 <= i {
			carry += int(a[i] - '0')
			i--
		}
		if 0 <= j {
			carry += int(b[j] - '0')
			j--
		}

		if carry%2 == 0 {
			res = append(res, '0')
		} else {
			res = append(res, '1')
		}
		carry /= 2
	}

	l, r := 0, len(res)-1
	for l < r {
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}
	return string(res)
}
