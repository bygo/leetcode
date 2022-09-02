package main

// https://leetcode.cn/problems/convert-a-number-to-hexadecimal

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	var buf []byte
	for idx := 7; 0 <= idx; idx-- {
		numHex := byte(num >> (idx * 4) & 0xf) // TODO
		if numHex == 0 && len(buf) == 0 {      // TODO
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
