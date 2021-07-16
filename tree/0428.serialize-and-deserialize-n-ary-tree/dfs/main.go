package main

import (
	"fmt"
	"strconv"
)

type Node struct {
	Val      int
	Children []*Node
}

//Title: Serialize and Deserialize N-ary Tree
// https://leetcode-cn.com/problems/serialize-and-deserialize-n-ary-tree

func main() {
	c := Codec{}
	s := c.serialize(&Node{
		Val: 1,
		Children: []*Node{
			{
				Val:      1,
				Children: nil,
			},
			{
				Val: 2,
				Children: []*Node{
					{
						Val:      3,
						Children: nil,
					},
					{
						Val: 4,
						Children: []*Node{
							{
								Val: 5,
								Children: []*Node{
									{
										Val:      6,
										Children: nil,
									},
									{
										Val:      7,
										Children: nil,
									},
								},
							},
						},
					},
				},
			},
		},
	})
	println(s)
	a := c.deserialize(s)
	//fmt.Printf("%+v", a)
	fmt.Printf("%+v", c.serialize(a))
}

type Codec struct {
}

func Constructor() *Codec {
	return &Codec{}
}

func (this *Codec) serialize(root *Node) string {
	if root == nil {
		return "[]"
	}
	var res = strconv.Itoa(root.Val) + ","
	if root.Children == nil {
		return res
	}
	res += "["
	for _, n := range root.Children {
		res += this.serialize(n)

	}
	return res + "]"
}

func (this *Codec) deserialize(data string) *Node {
	if data == "" {
		return nil
	}
	var parent = &Node{}
	var res = parent
	var sum int
	var stack []*Node
	var node *Node
	for _, v := range data {
		switch true {
		case 48 <= v && v <= 57:
			sum = sum*10 + int(v-48)
		case v == ',':
			node = &Node{Val: sum}
			parent.Children = append(parent.Children, node)
			sum = 0
		case v == '[':
			stack = append(stack, parent)
			parent = node
		case v == ']':
			top := len(stack) - 1
			parent = stack[top]
			stack = stack[:top]
		}
	}
	return res.Children[0]
}
