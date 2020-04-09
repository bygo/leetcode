package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var s int
var prefixSum map[int]int

func pathSum(root *TreeNode, sum int) int {
	s = sum
	prefixSum = map[int]int{}
	prefixSum[0] = 1
	return dfs(root, 0)
}

func dfs(root *TreeNode, cur int) int {
	if root == nil {
		return 0
	}
	cur += root.Val
	res := prefixSum[cur-s] //负数索引忽略。刚好为0，或者刚好为前缀  则+1
	prefixSum[cur]++        //加入前缀
	res += dfs(root.Left, cur) + dfs(root.Right, cur)
	prefixSum[cur]-- //取消前缀
	return res
}
