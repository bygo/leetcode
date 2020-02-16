package main

func addStrings(num1 string, num2 string) string {
	var carry byte
	var res []byte
	i, j := len(num1), len(num2)
	for i > 0 || j > 0 || carry > 0 {
		if i > 0 {
			i--
			carry += num1[i] - '0'
		}
		if j > 0 {
			j--
			carry += num2[j] - '0'
		}
		res = append(res, (carry%10)+'0')
		carry /= 10
	}
	l, r := 0, len(res)-1
	for l < r { //反转
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}

	return string(res)
}
