package main

// https://leetcode-cn.com/problems/convert-a-number-to-hexadecimal

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	var buf []byte
	for pos := 7; 0 <= pos; pos-- {
		numHex := byte(num >> (pos * 4) & 0xf) // 每4位1 代表1个16进制， 0~15
		if numHex == 0 && len(buf) == 0 {
			continue
		}
		if numHex <= 9 {
			buf = append(buf, numHex+'0')
		} else if 10 <= numHex {
			// 超过10 字母代替
			buf = append(buf, numHex+'a'-10)
		}
	}
	return string(buf)
}
