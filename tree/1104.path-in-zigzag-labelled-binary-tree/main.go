package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//Title: Path In Zigzag Labelled Binary Tree
//Link: https://leetcode-cn.com/problems/path-in-zigzag-labelled-binary-tree

func pathInZigZagTree(label int) []int {
	var res []int
	for 0 < label {
		res = append(res, label)
		label >>= 1
	}

	i, j := 0, len(res)-1
	for i < j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}

	y := len(res) % 2
	for k, v := range res {
		if k%2 == y {
			res[k] = 1<<(k+1) + 1<<k - v - 1
		}
	}

	return res
}
