# 前序遍历

### 递归

```
func preorder(root *TreeNode) {
	if root != nil {
		println(root.Val)    //前序遍历
		preorder(root.Left)  //左节点 遍历
		preorder(root.Right) //右节点 遍历
	}
}
```

### 迭代

```
func preorderStack(root *TreeNode) []int {
	var res []int
	var stack = []*TreeNode{root}
	for 0 < len(stack) {
		top := len(stack) - 1 //栈顶
		root = stack[top]     //出栈
		stack = stack[:top]
		if root != nil {
			res = append(res, root.Val) //根左右
			stack = append(stack, root.Right, root.Left)
		}
	}
	return res
}
```

### Morris 1

```
func preorderMorrisBreak(root *TreeNode) []int {
	var res []int
	var max *TreeNode
	for root != nil {
		res = append(res, root.Val) //前序遍历
		if root.Left == nil {
			root = root.Right //移动
		} else {
			max = root.Left //找左树最大节点
			for max.Right != nil {
				max = max.Right
			}

			//前序指针处理
			max.Right = root.Right
			root, root.Left = root.Left, nil //转为以Right为方向的单向树
		}
	}
	return res
}
```

### Morris 2

```
func preorderMorrisKeep(root *TreeNode) []int {
	var res []int
	var max *TreeNode
	for root != nil {
		if root.Left == nil {
			res = append(res, root.Val) //左节点为空 前序遍历
			root = root.Right           //移动到右节点
		} else {
			max = root.Left //寻找左树最大值
			for max.Right != nil && max.Right != root.Right {
				max = max.Right
			}

			//前序指针处理
			if max.Right == nil {
				res = append(res, root.Val) //前序遍历
				max.Right = root.Right      //指向
				root = root.Left            //移动到左节点
			} else { //已指向
				root = root.Right //跳跃
				max.Right = nil   //删指向
			}
		}
	}
	return res
}
```