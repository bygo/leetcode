package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode-cn.com/problems/next-greater-node-in-linked-list

func nextLargerNodes(head *ListNode) []int {
	var arr []int
	cur := head
	for cur != nil {
		arr = append(arr, cur.Val)
		cur = cur.Next
	}

	res := make([]int, len(arr))
	stack := make([]int, 0) // 单调递减栈
	for k := range arr {
		for {
			top := len(stack) - 1
			if top == -1 || arr[k] <= arr[stack[top]] { // 为空 或者 已经小于等于栈顶
				break
			}

			res[stack[top]] = arr[k] // 索引拿到它想要的更大值
			stack = stack[:top]      // 出栈
		}
		stack = append(stack, k) // 当前索引存起来 加入栈
	}
	return res
}
