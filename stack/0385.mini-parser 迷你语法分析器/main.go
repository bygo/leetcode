package main

//Link: https://leetcode-cn.com/problems/mini-parser

type NestedInteger struct {
}

func (n NestedInteger) IsInteger() bool { return true }

func (n NestedInteger) GetInteger() int { return 0 }

func (n NestedInteger) SetInteger(value int) {}

func (n NestedInteger) Add(elem NestedInteger) {}

func (n NestedInteger) GetList() []NestedInteger { return []NestedInteger{} }

func deserialize(s string) *NestedInteger {
	var stack = []NestedInteger{}
	n := len(s)
	i := 0
	var sign = 1
	for i < n {
		var num = 0
		switch s[i] {
		case ',':
			i++
		case '[':
			i++
			stack = append(stack, NestedInteger{})
		case ']':
			i++
			top := len(stack) - 1
			if 0 < top {
				stack[top-1].Add(stack[top])
				stack = stack[:top]
			}
		case '-':
			i++
			sign = -1
		default:
			for i < n && s[i] != ',' && s[i] != ']' {
				num = num*10 + int(s[i]-'0')
				i++
			}
			top := len(stack) - 1
			var ni NestedInteger
			ni.SetInteger(num * sign)
			if top == -1 {
				return &ni
			} else {
				stack[top].Add(ni)
			}
			sign = 1
		}
	}
	return &stack[0]
}
