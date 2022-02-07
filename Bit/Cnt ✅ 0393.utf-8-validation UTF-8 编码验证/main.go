package main

// https://leetcode-cn.com/problems/utf-8-validation

// ❓ UTF-8 编码验证

func validUtf8(data []int) bool {
	cnt := 0
	for _, num := range data {
		// 首段 验证
		if cnt == 0 {
			if num>>5 == 0b110 {
				cnt = 1
			} else if num>>4 == 0b1110 {
				cnt = 2
			} else if num>>3 == 0b11110 {
				cnt = 3
			} else if num>>7 != 0 {
				return false
			}
		} else {
			// 其他段 验证
			if num>>6 != 0b10 {
				return false
			}
			cnt--
		}
	}
	return cnt == 0
}
