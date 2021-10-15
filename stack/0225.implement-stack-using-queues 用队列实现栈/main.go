package main

// https://leetcode-cn.com/problems/implement-stack-using-queues

type MyStack struct {
	queue []int
}

func Constructor() (s MyStack) {
	return MyStack{}
}

// 要嘛push o(n)  要嘛pop o(n)
func (s *MyStack) Push(x int) {
	n := len(s.queue)
	s.queue = append(s.queue, x)
	s.queue = append(s.queue, s.queue[:n]...)
	s.queue = s.queue[n:]
}

func (s *MyStack) Pop() int {
	v := s.queue[0]
	s.queue = s.queue[1:]
	return v
}

func (s *MyStack) Top() int {
	return s.queue[0]
}

func (s *MyStack) Empty() bool {
	return len(s.queue) == 0
}


