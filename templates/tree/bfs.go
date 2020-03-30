package tree

func levelOrder(root *TreeNode) {
	var queue = []*TreeNode{root}
	var level int
	for 0 < len(queue) {
		length := len(queue)
		for 0 < length {
			length--
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			println(queue[0].Val) //当前节点值
			println(level)        //层级
			queue = queue[1:]
		}
		level++
	}
}
