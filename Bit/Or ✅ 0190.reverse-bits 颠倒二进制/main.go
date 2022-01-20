package main

// https://leetcode-cn.com/problems/reverse-bits

// ❓ 颠倒二进制

func reverseBits(n uint32) uint32 {
	var num uint32
	for pos := 31; 0 <= pos; pos-- {
		num |= n & 1 << pos
		n >>= 1
	}
	return num
}
