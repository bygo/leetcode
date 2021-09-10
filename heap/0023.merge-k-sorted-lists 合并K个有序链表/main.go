package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type heap []*ListNode

func mergeKLists(lists []*ListNode) *ListNode {
	k := len(lists)
	h := &heap{}
	for i := 0; i < k; i++ {
		if lists[i] != nil {
			h.Push(lists[i])
		}
	}

	zero := new(ListNode)
	pre := zero
	for 0 < len(*h) {
		tmp := h.Pop()
		if tmp.Next != nil {
			h.Push(tmp.Next)
		}

		pre.Next = tmp
		pre = pre.Next
	}

	return zero.Next
}

func (h *heap) Pop() *ListNode {
	n := len(*h) - 1
	h.Swap(0, n)
	h.down(0)
	x := (*h)[n]
	*h = (*h)[:n]
	return x
}

func (h *heap) Init() {
	n := len(*h)
	for i := n/2 - 1; 0 <= i; i-- {
		h.down(i)
	}
}
func (h *heap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *heap) Push(x *ListNode) {
	*h = append(*h, x)
	h.up(len(*h) - 1)
}

func (h heap) up(j int) {
	for {
		i := (j - 1) / 2
		if i == j || !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		j = i
	}
}

func (h heap) down(i int) bool {
	raw := i
	n := len(h) - 1
	for {
		l := 2*i + 1
		if l < 0 || n <= l {
			break
		}
		if r := l + 1; r < n && h.Less(r, l) {
			l = r
		}
		if !h.Less(l, i) {
			break
		}
		h.Swap(i, l)
		i = l
	}
	return raw < i
}

func (h heap) Fix(i int) {
	if !h.down(i) {
		h.up(i)
	}
}

func (h heap) Remove(i int) *ListNode {
	n := len(h) - 1
	if n != i {
		h.Swap(i, n)
		if !h.down(i) {
			h.up(i)
		}
	}
	return h.Pop()
}

func (h heap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}
