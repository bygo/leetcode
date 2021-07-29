package main

// https://leetcode-cn.com/problems/add-strings/

// 1:44
func addStrings(num1 string, num2 string) string {
	var i = len(num1)
	var j = len(num2)
	var carry byte
	var res []byte

	for 0 < i || 0 < j || 0 < carry {
		if 0 < i {
			i--
			carry += num1[i] - '0'
		}

		if 0 < j {
			j--
			carry += num2[j] - '0'
		}

		res = append(res, carry%10+'0')
		carry /= 10
	}

	l, r := 0, len(res)-1
	for l < r {
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}
	return string(res)
}
