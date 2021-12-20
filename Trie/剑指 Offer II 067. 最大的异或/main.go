package main

// https://leetcode-cn.com/problems/ms70jA

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
	var x int
	var cur = t
	for i := h; 0 <= i; i-- {
		bit := num >> i & 1
		if bit == 0 {
			if cur.right != nil {
				x = x*2 + 1
				cur = cur.right
			} else {
				x = x * 2
				cur = cur.left
			}
		} else if bit == 1 {
			if cur.left != nil {
				x = x*2 + 1
				cur = cur.left
			} else {
				x = x * 2
				cur = cur.right
			}
		}
	}
	return x
}

func findMaximumXOR(nums []int) (x int) {
	root := &trie{}
	var res int
	for _, num := range nums {
		root.add(num)
		cur := root.check(num)
		if res < cur {
			res = cur
		}
	}
	return res
}
