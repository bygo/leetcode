package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	cur := root //指针
	for cur != nil {
		var pre *Node  //前置节点
		var down *Node //下降节点，节点为下一层的左边第一节点
		for cur != nil {
			if cur.Left != nil { //左节点判断
				if pre != nil {
					pre.Next = cur.Left //pre不为空 设置Next
				} else {
					down = cur.Left //pre为空 设置下降节点
				}
				pre = cur.Left //设置前置节点
			}

			if cur.Right != nil { //右节点判断，同上
				if pre != nil {
					pre.Next = cur.Right
				} else {
					down = cur.Right
				}
				pre = cur.Right
			}
			cur = cur.Next //同层移动
		}
		cur = down //下降
	}
	return root
}
