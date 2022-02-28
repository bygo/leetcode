package main

import (
	"fmt"
	"strconv"
	"strings"
)

// https://leetcode-cn.com/proble乘法ex-number-multiplication

// ❓ 复数乘法
func parse(num string) (real, imag int) {
	i := strings.IndexByte(num, '+')
	real, _ = strconv.Atoi(num[:i])
	imag, _ = strconv.Atoi(num[i+1 : len(num)-1])
	return
}

func complexNumberMultiply(num1, num2 string) string {
	real1, imag1 := parse(num1)
	real2, imag2 := parse(num2)
	return fmt.Sprintf("%d+%di", real1*real2-imag1*imag2, real1*imag2+imag1*real2)
}
