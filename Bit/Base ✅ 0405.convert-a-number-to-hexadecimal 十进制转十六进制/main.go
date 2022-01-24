package main

// https://leetcode-cn.com/problems/convert-a-number-to-hexadecimal

// ❓十进制转十六进制

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	var buf []byte
	for pos := 7; 0 <= pos; pos-- {
		numHex := byte(num >> (pos * 4) & 0xf) // 每4位1 代表1个16进制， 15
		if 0 < numHex || 0 < len(buf) {
			if numHex <= 9 {
				numHex = numHex + '0'
			} else {
				numHex = numHex - 10 + 'a'
			}
			buf = append(buf, numHex)
		}
	}
	return string(buf)
}
