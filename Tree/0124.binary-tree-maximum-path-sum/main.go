package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res int

func maxPathSum(root *TreeNode) (sum int) {
	if root == nil {
		return 0
	}
	res = -1 << 63 //用来判断 所有节点都是负数时，应返回哪一个。
	dfs(root)
	return res
}

func dfs(root *TreeNode) (sum int) {
	if root == nil {
		return
	}
	lMax := max(0, dfs(root.Left))     //左分支 最大贡献
	rMax := max(0, dfs(root.Right))    //右分支 最大贡献
	res = max(res, root.Val+lMax+rMax) //与当前结果对比
	return root.Val + max(lMax, rMax)  //返回单分支最大值
}

func max(nums ...int) int {
	max := -1 << 63
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
}
