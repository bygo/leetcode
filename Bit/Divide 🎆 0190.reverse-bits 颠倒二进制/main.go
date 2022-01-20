package main

// https://leetcode-cn.com/problems/reverse-bits

// ❓ 颠倒二进制

const (
	m1 = 0b01010101010101010101010101010101
	m2 = 0b00110011001100110011001100110011
	m4 = 0b00001111000011110000111100001111
	m8 = 0b00000000111111110000000011111111
)

func reverseBits(n uint32) uint32 {
	n = n>>1&m1 | n&m1<<1 // n>>1 左移右 ， n&m1 右合法 移左
	n = n>>2&m2 | n&m2<<2
	n = n>>4&m4 | n&m4<<4
	n = n>>8&m8 | n&m8<<8
	return n>>16 | n<<16
}
