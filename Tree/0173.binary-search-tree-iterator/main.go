package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	root *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{root: root}
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	root := this.n(true)
	return root.Val
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	root := this.n(false)
	return root != nil
}

func (this *BSTIterator) n(move bool) *TreeNode {
	var max *TreeNode

	for this.root != nil {
		if this.root.Left == nil {
			if move {
				defer func() {
					this.root = this.root.Right
				}()
			}
			return this.root
		} else {
			//寻找左树最大节点
			max = this.root.Left
			for max.Right != nil {
				max = max.Right
			}

			//最大节点指向root
			max.Right = this.root

			//root = root.Left :移动root到root.left
			//root.Left = nil  :砍左子树，避免下一次遍历到root时，再进入到左子树
			this.root, this.root.Left = this.root.Left, nil
		}
	}
	return nil
}
