package main

func main() {
	h := heap{9, 3, 6, 4, 5, 7, 8, 1, 2, 4, 3, 2, 1, 3, 6, 7, 8}
	h.Init()
	for 0 < len(h) {
		println(h.Pop())
	}
}

type heap []int

func (h *heap) Init() {
	l := len(*h)
	i := l/2 - 1
	for 0 <= i {
		h.down(i)
		i--
	}
}
func (h *heap) Pop() int {
	n := len(*h) - 1
	h.Swap(0, n)
	h.down(0)
	x := (*h)[n]
	*h = (*h)[:n]
	return x
}

func (h *heap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *heap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h heap) up(cur int) {
	for {
		parent := (cur - 1) / 2
		if parent == cur || h.Less(cur, parent) {
			break
		}
		h.Swap(parent, cur)
		cur = parent
	}
}

func (h heap) down(parent int) bool {
	n := len(h) - 1
	raw := parent
	for {
		l := 2*parent + 1
		if n <= l { // 超过 length
			break
		}
		r := l + 1
		if r < n && h.Less(l, r) {
			l = r
		}
		if h.Less(l, parent) {
			break
		}
		h.Swap(parent, l)
		parent = l
	}
	return raw != parent
}

func (h heap) Fix(i int) {
	if !h.down(i) {
		h.up(i)
	}
}

func (h heap) Remove(i int) int {
	n := len(h) - 1
	if n != i {
		h.Swap(i, n)
		h.Fix(i)
	}
	return h.Pop()
}
