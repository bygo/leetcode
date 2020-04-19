package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res []int

func findTarget(root *TreeNode, k int) bool {
	res = []int{}
	dfs(root)
	return twoSum(res,k) != nil
}

func dfs(root *TreeNode) {
	if root != nil {
		dfs(root.Left)
		res = append(res, root.Val)
		dfs(root.Right)
	}
}

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int, len(nums))
	for rawK, rawV := range nums {
		if hashK, ok := hashMap[target-rawV]; ok {
			return []int{rawK, hashK}
		}
		hashMap[rawV] = rawK
	}
	return nil
}
