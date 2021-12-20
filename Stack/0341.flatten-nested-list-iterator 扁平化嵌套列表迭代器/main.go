package main

// https://leetcode-cn.com/problems/flatten-nested-list-iterator

type NestedInteger struct {
}

func (n NestedInteger) IsInteger() bool { return true }

func (n NestedInteger) GetInteger() int { return 0 }

func (n NestedInteger) SetInteger(value int) {}

func (n NestedInteger) Add(elem NestedInteger) {}

func (n NestedInteger) GetList() []*NestedInteger { return []*NestedInteger{} }

// 先迭代
type NestedIterator2 struct {
	vals []int
}

func Constructor2(nestedList []*NestedInteger) *NestedIterator {
	var vals []int
	dfs(nestedList, &vals)
	return &NestedIterator{vals}
}

func (it *NestedIterator2) Next() int {
	val := it.vals[0]
	it.vals = it.vals[1:]
	return val
}

func (it *NestedIterator2) HasNext() bool {
	return len(it.vals) != 0
}

func dfs(nestedList []*NestedInteger, vals *[]int) {
	for _, nest := range nestedList {
		if nest.IsInteger() {
			*vals = append(*vals, nest.GetInteger())
		} else {
			dfs(nest.GetList(), vals)
		}
	}
}

// 栈
type NestedIterator struct {
	stack [][]*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{[][]*NestedInteger{nestedList}}
}

func (it *NestedIterator) Next() int {
	top := len(it.stack) - 1
	queue := it.stack[top]
	val := queue[0].GetInteger()
	it.stack[top] = queue[1:]
	return val
}

func (it *NestedIterator) HasNext() bool {
	for 0 < len(it.stack) {
		top := len(it.stack) - 1
		queue := it.stack[top]
		if len(queue) == 0 {
			it.stack = it.stack[:top]
			continue
		}
		nest := queue[0]
		if nest.IsInteger() {
			return true
		}
		it.stack[top] = queue[1:]
		it.stack = append(it.stack, nest.GetList())
	}
	return false
}
