package main

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	l1, l2 := len(num1), len(num2)
	l := l1 + l2
	carry := make([]byte, l)
	for i := l1; i > 0; i-- {
		index := l2 + i - 1
		n1 := num1[i-1] - '0'
		for j := l2; j > 0; j-- {
			n := carry[index] + n1*(num2[j-1]-'0') //计算
			carry[index] = n % 10                  //当前值
			index--
			carry[index] += n / 10 //进位
		}
	}

	j := -1
	for i := 0; i < l; i++ {
		if carry[i] != 0 && j == -1 {
			j = i
		}
		carry[i] += '0'
	}
	return string(carry[j:])
}
