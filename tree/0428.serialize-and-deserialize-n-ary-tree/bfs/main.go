package main

import (
	"strconv"
	"strings"
)

type Node struct {
	Val      int
	Children []*Node
}

//Title: Serialize and Deserialize N-ary Tree
//Link: https://leetcode-cn.com/problems/serialize-and-deserialize-n-ary-tree

type Codec struct {
}

func Constructor() *Codec {
	return &Codec{}
}

func (this *Codec) serialize(root *Node) string {
	if root == nil {
		return ""
	}

	var queue = []*Node{root}
	var res = []string{strconv.Itoa(root.Val)}
	for {
		counter := len(queue)
		if counter == 0 {
			break
		}
		res = append(res, "d")

		for qk, q := range queue[:counter] {
			if qk != 0 {
				res = append(res, "#")
			}
			for ck, c := range q.Children {
				if ck != 0 {
					res = append(res, ",")
				}
				queue = append(queue, c)
				res = append(res, strconv.Itoa(c.Val))
			}
		}
		queue = queue[counter:]
	}

	return strings.Join(res, "")
}

func (this *Codec) deserialize(data string) *Node {
	if data == "" {
		return nil
	}
	var pre, parent = &Node{}, &Node{}
	var queue = []*Node{pre}
	for _, d := range strings.Split(data, "d") {
		counter := len(queue)
		for k, n := range strings.Split(d, "#") {
			parent = queue[k]
			for _, q := range strings.Split(n, ",") {
				if q != "" {
					i, _ := strconv.Atoi(q)
					c := &Node{Val: i}
					parent.Children = append(parent.Children, c)
					queue = append(queue, c)
				}
			}
		}
		queue = queue[counter:]
	}

	return pre.Children[0]
}
