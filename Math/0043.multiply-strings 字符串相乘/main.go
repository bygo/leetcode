package main

import "fmt"

// https://leetcode.cn/problems/multiply-strings/

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

func main() {
	s := "2393706880236110407059624696967828762752651982730115221690437821508229419410771541532394006597463715513741725852432559057224478815116557380260390432211227579663571046845842281704281749571110076974264971989893607137140456254346955633455446057823738757323149856858154529105301197388177242583658641529908583934918768953462557716z97438020429952944646288084173334701047574188936201324845149110176716130267041674438237608038734431519439828191344238609567530399189316846359766256507371240530620697102864238792350289978450509162697068948604722646739174590530336510475061521094503850598453536706982695212493902968251702853203929616930291257062173c79487281900662343830648295410"
	var res = map[string]int{}
	var cur []byte
	for i := range s {
		ch := s[i]
		if '0' <= ch && ch <= '9' {
			cur = append(cur, ch)
		} else if 0 < len(cur) {
			println(string(ch))
			res[string(cur)]++
			cur = []byte{}
		}
	}
	if 0 < len(cur) {
		res[string(cur)]++
	}
	fmt.Printf("%+v", len(res))
	//println(math.Pow(2, -2) == pow(2, -2))
	//println(math.Pow(312, -31) == pow(312, -31))
	//println(math.Pow(-3110.2233, -31) == pow(-3110.2233, -31))
}
