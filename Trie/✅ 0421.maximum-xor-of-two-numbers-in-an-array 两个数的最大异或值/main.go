package main

// https://leetcode.cn/problems/maximum-xor-of-two-numbers-in-an-array

// ❓ 两个数的最大异或值
// ⚠️ 贪心路径

type trie struct {
	bit [2]*trie
}

func (root *trie) add(num int) {
	cur := root
	for i := 30; 0 <= i; i-- {
		bit := num >> i & 1
		if bit == 0 {
			if cur.bit[0] == nil {
				cur.bit[0] = &trie{}
			}
			cur = cur.bit[0]
		} else if bit == 1 {
			if cur.bit[1] == nil {
				cur.bit[1] = &trie{}
			}
			cur = cur.bit[1]
		}
	}
}

func (root *trie) xor(num int) int {
	var numMax int
	var node = root
	// 从大到小贪心路径
	for i := 30; 0 <= i; i-- {
		bit := num >> i & 1
		switch bit {
		case 0:
			if node.bit[1] != nil {
				numMax = numMax*2 + 1
				node = node.bit[1]
			} else {
				numMax = numMax * 2
				node = node.bit[0]
			}
		case 1:
			if node.bit[0] != nil {
				numMax = numMax*2 + 1
				node = node.bit[0]
			} else {
				numMax = numMax * 2
				node = node.bit[1]
			}
		}
	}
	return numMax
}

func findMaximumXOR(nums []int) int {
	root := &trie{}
	var numMax int
	// 构建一条路径,避免panic
	root.add(nums[0])
	numsL := len(nums)
	for i := 1; i < numsL; i++ {
		num := nums[i]
		numCur := root.xor(num)
		if numMax < numCur {
			numMax = numCur
		}
		root.add(num)
	}
	return numMax
}
