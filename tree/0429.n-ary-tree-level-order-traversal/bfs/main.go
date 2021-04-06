package main

type Node struct {
	Val      int
	Children []*Node
}

var res [][]int

func levelOrder(root *Node) [][]int {
	res = [][]int{}
	if root == nil {
		return res
	}

	var level int
	var queue = []*Node{root}
	for {
		counter := len(queue)
		if counter == 0 {
			break
		}

		res = append(res, []int{})
		for _, v := range queue[:counter] {
			res[level] = append(res[level], v.Val)
			for _, c := range v.Children {
				queue = append(queue, c)
			}
		}
		level++
		queue = queue[counter:]
	}
	return res
}
