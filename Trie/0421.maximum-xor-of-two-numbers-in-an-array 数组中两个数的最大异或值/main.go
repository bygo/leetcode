package main

// https://leetcode-cn.com/problems/maximum-xor-of-two-numbers-in-an-array

const h = 30

type trie struct {
	left, right *trie
}

func (t *trie) add(num int) {
	cur := t
	for i := h; 0 <= i; i-- {
		bit := num >> i & 1
		if bit == 0 {
			if cur.left == nil {
				cur.left = &trie{}
			}
			cur = cur.left
		} else if bit == 1 {
			if cur.right == nil {
				cur.right = &trie{}
			}
			cur = cur.right
		}
	}
}

func (t *trie) check(num int) int {
	var res int
	var cur = t
	for i := h; 0 <= i && cur != nil; i-- {
		bit := num >> i & 1
		res <<= 1
		if bit == 0 {
			if cur.right != nil {
				res += 1
				cur = cur.right
			} else {
				cur = cur.left
			}
		} else if bit == 1 {
			if cur.left != nil {
				res += 1
				cur = cur.left
			} else {
				cur = cur.right
			}
		}
	}
	return res
}

func findMaximumXOR(nums []int) (x int) {
	root := &trie{}
	var res int
	for _, num := range nums {
		cur := root.check(num)
		if res < cur {
			res = cur
		}
		root.add(num)
	}
	return res
}
