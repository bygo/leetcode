package main

// https://leetcode-cn.com/problems/map-sum-pairs

// ❓ 前缀为 prefix 的 key 的 值的总和。

type MapSum struct {
	chMp [26]*MapSum
	val  int
}

func Constructor() MapSum {
	return MapSum{
		chMp: [26]*MapSum{},
	}
}

func (root *MapSum) Insert(key string, val int) {
	node := root
	for i := range key {
		ch := key[i] - 'a'
		if node.chMp[ch] == nil {
			node.chMp[ch] = &MapSum{
				chMp: [26]*MapSum{},
			}
		}
		node = node.chMp[ch]
	}
	node.val = val
}

func (root *MapSum) Sum(prefix string) int {
	node := root
	var sum int
	for i := range prefix {
		ch := prefix[i] - 'a'
		if node.chMp[ch] == nil {
			return sum
		}
		node = node.chMp[ch]
	}

	var que = []*MapSum{node}
	for {
		cnt := len(que)
		if cnt == 0 {
			break
		}
		for i := range que[:cnt] {
			nodeCur := que[i]
			sum += nodeCur.val
			for ch := range nodeCur.chMp {
				elem := nodeCur.chMp[ch]
				if elem != nil {
					que = append(que, elem)
				}
			}
		}
		que = que[cnt:]
	}
	return sum
}
