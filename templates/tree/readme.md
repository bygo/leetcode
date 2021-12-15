# 前序遍历

### 递归

```go
func preorder(root *TreeNode) {
	if root != nil {
		println(root.Val)    //前序遍历
		preorder(root.Left)  //左节点 遍历
		preorder(root.Right) //右节点 遍历
	}
}
```

### 迭代

```go
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

```go
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

```go
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

# 中序遍历

### 递归

```go
func inorder(root *TreeNode) {
	if root != nil {
		inorder(root.Left)  //左节点 遍历
		println(root.Val)   //中序遍历
		inorder(root.Right) //右节点 遍历
	}
}
```

### 迭代

```go
func inorderStack(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode

	for 0 < len(stack) || root != nil { //root != nil 只为了第一次root判断，必须放最后
		for root != nil {
			stack = append(stack, root) //入栈
			root = root.Left            //移至最左
		}

		top := len(stack) - 1             //栈顶
		res = append(res, stack[top].Val) //中序遍历
		root = stack[top].Right           //右节点会进入for循环，如果=nil，继续出栈
		stack = stack[:top]               //出栈
	}
	return res
}
```

### Morris 1

```go
func inorderMorrisBreak(root *TreeNode) []int {
	var res []int
	var max *TreeNode

	for root != nil {
		if root.Left == nil {
			res = append(res, root.Val) //中序遍历
			root = root.Right           //移动
		} else {
			max = root.Left //寻找左树最大值
			for max.Right != nil {
				max = max.Right
			}

			//中序指针处理
			max.Right = root                 //左树最大值连接root
			root, root.Left = root.Left, nil //移动左节点，砍左树
		}
	}
	return res
}
```

### Morris 2

```go
func inorderMorrisKeep(root *TreeNode) []int {
	var res []int
	var max *TreeNode
	for root != nil {
		if root.Left == nil {
			res = append(res, root.Val) //中序遍历
			root = root.Right
		} else {
			max = root.Left //寻找左树最大值
			for max.Right != nil && max.Right != root {
				max = max.Right
			}

			//中序指针处理
			if max.Right == nil {
				max.Right = root //左树最大值指向 root
				root = root.Left //移动到左节点
			} else { //max.Right 指向root
				res = append(res, root.Val) //中序遍历
				root = root.Right           //直接跳跃到root.Right
				max.Right = nil             //删指向
			}
		}
	}
	return res
}
```

# 后序遍历

### 递归

```go
func postorder(root *TreeNode) {
	if root != nil {
		postorder(root.Left)
		postorder(root.Right)
		println(root.Val)
	}
}
```

### 迭代

```go
func postorderStack(root *TreeNode) []int {
	var res []int
	var stack = []*TreeNode{root}
	for 0 < len(stack) {
		top := len(stack) - 1 //栈顶
		root = stack[top]     //出栈
		stack = stack[:top]
		if root != nil {
			res = append(res, root.Val) //根右左
			stack = append(stack, root.Left, root.Right)
		}
	}

	//反转 变成后序遍历 左右根
	l, r := 0, len(res)-1
	for l < r {
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}
	return res
}
```

### Morris 1

```go

func postorderMorrisBreak(root *TreeNode) []int {
	var res []int
	var min *TreeNode
	for root != nil {
		res = append(res, root.Val) //后序遍历
		if root.Right == nil {
			root = root.Left //移动
		} else {
			min = root.Right //寻找右树最小节点
			for min.Left != nil && min.Left != root {
				min = min.Left
			}

			//后序指针处理
			min.Left = root.Left
			root, root.Right = root.Right, nil
		}
	}

	l, r := 0, len(res)-1
	for l < r {
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}
	return res
}
```

### Morris 2

```go
func postorderMorrisKeep(root *TreeNode) []int {
	var res []int
	var min *TreeNode
	for root != nil {
		if root.Right == nil {
			res = append(res, root.Val) //后序遍历
			root = root.Left            //移动
		} else {
			min = root.Right //寻找右树最小节点
			for min.Left != nil && min.Left != root.Left {
				min = min.Left
			}

			if min.Left == nil {
				res = append(res, root.Val) //后序遍历
				min.Left = root.Left
				root = root.Right
			} else {
				root = root.Left //跳跃
				min.Left = nil
			}
		}
	}

	l, r := 0, len(res)-1
	for l < r {
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}
	return res
}
```

# 层次遍历

```go
func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	var queue = []*TreeNode{root}
	var level int //层级
	for {
		counter := len(queue)
		if counter == 0 {
			break
		}
		res = append(res, []int{})
		for 0 < counter {
			counter--
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			res[level] = append(res[level], queue[0].Val)
			queue = queue[1:]
		}
		level++
	}
	return res
}
```