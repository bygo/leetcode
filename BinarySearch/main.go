package main

func baseNeg2(n int) string {
	if n == 0 {
		return "0"
	}
	const BASE = -2
	var res []byte
	for n != 1 {
		if n > 0 {
			res = append(res, byte('0'+n%BASE))
			n /= BASE
		} else {
			cur := n % BASE
			if cur == -1 {
				n = n/BASE + 1
				res = append(res, '0'+byte(1))
			} else {
				res = append(res, '0'+byte(cur))
				n /= BASE
			}
		}
	}
	res = append(res, '0'+byte(1))

	length := len(res)

	for i := 0; i < length/2; i++ {
		res[i], res[length-1-i] = res[length-1-i], res[i]
	}
	return string(res)
}
