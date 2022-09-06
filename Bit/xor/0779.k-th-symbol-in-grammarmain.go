package main

import "math/bits"

// https://leetcode.cn/problems/k-th-symbol-in-grammar

/**
					0
			 0   			1
		0       1        1       0
     0   1	  1   0	   1   0   0   1
*/

func kthGrammar(n int, k int) int {
	if n == 1 {
		return 0
	}
	// 奇数`子节点`原值，即0，不改变原值
	// 偶数取反,即1，改变原值 1^1=0 ,1^0=1
	return (k+1)%2 ^ kthGrammar(n-1, (k+1)/2) // TODO 父节点
}

// mirror
func kthGrammar(n int, k int) int {
	if n == 1 {
		return 0
	}
	half := 1 << (n - 2)
	if k <= half {
		return kthGrammar(n-1, k) // TODO 前一半原值
	}
	return kthGrammar(n-1, k-half) ^ 1 // TODO 取反
}

// flip
func kthGrammar(n int, k int) int {
	return bits.OnesCount(uint(k-1)) % 2 // TODO 0 ~ k-1
	// 选择1：左节点为1，右节点为0
	// 选择0：左节点为0，右节点为1
	// 每个数字 代表第几个索引，选择了几次1
}
