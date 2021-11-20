package main

// https://leetcode-cn.com/problems/map-sum-pairs

type MapSum struct {
	child map[byte]*MapSum
	Val   int
}

func Constructor() MapSum {
	return MapSum{
		child: map[byte]*MapSum{},
	}
}

func (ms *MapSum) Insert(key string, val int) {
	node := ms
	for i := range key {
		ch := key[i]
		if node.child[ch] == nil {
			node.child[ch] = &MapSum{
				child: map[byte]*MapSum{},
			}
		}
		node = node.child[ch]
	}
	node.Val = val
}

func (ms *MapSum) Sum(prefix string) int {
	node := ms
	var sum int
	for i := range prefix {
		ch := prefix[i]
		if node.child[ch] == nil {
			return sum
		}
		node = node.child[ch]
	}

	var queue = []*MapSum{node}
	for {
		cnt := len(queue)
		if cnt == 0 {
			break
		}
		for i := range queue {
			sum += queue[i].Val
			for b := range queue[i].child {
				queue = append(queue, queue[i].child[b])
			}
		}
		queue = queue[cnt:]
	}
	return sum
}
