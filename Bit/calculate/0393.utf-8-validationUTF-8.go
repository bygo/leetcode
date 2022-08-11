package main

// https://leetcode.cn/problems/utf-8-validation

// ❓ UTF-8 编码验证

func validUtf8(bytes []int) bool {
	cnt := 0
	for _, num := range bytes {
		// 首段 验证
		if cnt == 0 {
			if num>>5 == 0b110 {
				// 2 字符
				cnt = 1
			} else if num>>4 == 0b1110 {
				// 3 字符
				cnt = 2
			} else if num>>3 == 0b11110 {
				// 4 字符
				cnt = 3
			} else if num>>7 != 0 {
				// 不为0
				return false
			}
		} else {
			// 其他剩余段 验证
			if num>>6 != 0b10 {
				return false
			}
			cnt--
		}
	}
	return cnt == 0
}
