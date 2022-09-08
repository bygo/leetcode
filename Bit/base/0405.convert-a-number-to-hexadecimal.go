package main

// https://leetcode.cn/problems/convert-a-number-to-hexadecimal

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	var buf []byte
	for idx := 28; 0 <= idx; idx -= 4 {
		numHex := byte(num >> idx & 0xf)  // TODO 取有效位
		if numHex == 0 && len(buf) == 0 { // TODO 为0忽略
			continue
		}
		if numHex <= 9 {
			buf = append(buf, numHex+'0')
		} else if 10 <= numHex {
			buf = append(buf, numHex+'a'-10)
		}
	}
	return string(buf)
}
