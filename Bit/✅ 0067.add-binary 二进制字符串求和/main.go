package main

// https://leetcode-cn.com/problems/add-binary

// ❓  二进制字符串求和

func addBinary(a string, b string) string {
	carry := 0
	var aTop, bTop = len(a) - 1, len(b) - 1
	var num = []byte{}
	for 0 <= aTop || 0 <= bTop || carry != 0 {
		if 0 <= aTop {
			carry += int(a[aTop] - '0')
			aTop--
		}

		if 0 <= bTop {
			carry += int(b[bTop] - '0')
			bTop--
		}

		if carry%2 == 0 {
			// 0和2 为0
			num = append(num, '0')
		} else {
			// 1和3 为1
			num = append(num, '1')
		}
		// 2 3进1
		carry /= 2
	}

	left, right := 0, len(num)-1
	for left < right {
		num[left], num[right] = num[right], num[left]
		left++
		right--
	}
	return string(num)
}
