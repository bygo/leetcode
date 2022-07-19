package main

// https://leetcode.cn/problems/lru-cache

type LRUCache struct {
	cache      map[int]*Node
	head, tail *Node
	capacity   int
	size       int
}

type Node struct {
	key, val   int
	prev, next *Node
}

func Constructor(capacity int) LRUCache {
	head := New(0, 0)
	tail := New(0, 0)
	head.next = tail
	tail.prev = head
	return LRUCache{
		cache:    map[int]*Node{},
		head:     head,
		tail:     tail,
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	n, ok := this.cache[key]
	if !ok {
		return -1
	}
	this.moveHead(n)
	return n.val
}

func New(key, value int) *Node {
	return &Node{key: key, val: value}
}

func (this *LRUCache) Put(key int, value int) {
	n, ok := this.cache[key]
	if ok {
		n.val = value
		this.moveHead(n)
	} else {
		if this.capacity <= this.size {
			this.remove(this.tail.prev)
			this.size--
		}

		this.size++
		n := New(key, value)
		this.cache[n.key] = n
		n.next = this.head.next
		this.head.next.prev = n
		this.head.next = n
		n.prev = this.head
	}
}

func (this *LRUCache) remove(n *Node) {
	n.prev.next = n.next
	n.next.prev = n.prev
	delete(this.cache, n.key)
}

func (this *LRUCache) moveHead(n *Node) {
	n.prev.next = n.next
	n.next.prev = n.prev

	this.head.next.prev = n
	n.next = this.head.next
	n.prev = this.head
	this.head.next = n
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
