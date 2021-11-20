package main

import (
	"strconv"
	"strings"
)

type Node struct {
	Val      int
	Children []*Node
}

// serialize and Deserialize N-ary Tree
// https://leetcode-cn.com/problems/serialize-and-deserialize-n-ary-tree/

type Codec struct {
}

func Constructor() *Codec {
	return &Codec{}
}

func (c *Codec) serialize(root *Node) string {
	if root == nil {
		return ""
	}

	var queue = []*Node{root}
	var res = []string{strconv.Itoa(root.Val)}
	for 0 < len(queue) {
		var cnt = len(queue)
		res = append(res, "d")
		for i, p := range queue[:cnt] {
			if i != 0 {
				res = append(res, "#")
			}
			for j, c := range p.Children {
				if j != 0 {
					res = append(res, ",")
				}
				queue = append(queue, c)
				res = append(res, strconv.Itoa(c.Val))
			}
		}
		queue = queue[cnt:]
	}

	return strings.Join(res, "")
}

func (c *Codec) deserialize(data string) *Node {
	if data == "" {
		return nil
	}
	var pre, parent = &Node{}, &Node{}
	var queue = []*Node{pre}
	for _, d := range strings.Split(data, "d") {
		var cnt = len(queue)
		for k, n := range strings.Split(d, "#") {
			parent = queue[k]
			for _, q := range strings.Split(n, ",") {
				if q != "" {
					v, _ := strconv.Atoi(q)
					c := &Node{Val: v}
					parent.Children = append(parent.Children, c)
					queue = append(queue, c)
				}
			}
		}
		queue = queue[cnt:]
	}

	return pre.Children[0]
}
