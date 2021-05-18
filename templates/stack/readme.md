# 栈的特性
 
### [用栈实现队列](https://leetcode-cn.com/problems/implement-queue-using-stacks)

```go
type MyQueue struct {
	inStack, outStack []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (q *MyQueue) Push(x int) {
	q.inStack = append(q.inStack, x)
}

// 双栈 倒2次顺序
func (q *MyQueue) in2out() {
	for 0 < len(q.inStack) {
		q.outStack = append(q.outStack, q.inStack[len(q.inStack)-1])
		q.inStack = q.inStack[:len(q.inStack)-1]
	}
}

func (q *MyQueue) Pop() int {
	if len(q.outStack) == 0 {
		q.in2out()
	}
	x := q.outStack[len(q.outStack)-1]
	q.outStack = q.outStack[:len(q.outStack)-1]
	return x
}
```

###  [用队列实现栈](https://leetcode-cn.com/problems/implement-stack-using-queues)

```go
type MyStack struct {
	queue []int
}

func Constructor() (s MyStack) {
	return MyStack{}
}

// 每次都是 o(n)
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
```

### [从尾到头打印链表](https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof)

```go
func reversePrint(head *ListNode) []int {
	var res = []int{}
	var stack []*ListNode
	for head != nil {
		stack = append(stack, head)
		head = head.Next
	}

	for {
		top := len(stack) - 1
		if -1 == top {
			break
		}
		res = append(res, stack[top].Val)
		stack = stack[:top]
	}
	return res
}
```


# 单调栈


### [接雨水](https://leetcode-cn.com/problems/trapping-rain-water)

```go
func trap(height []int) int {
	var res = 0
	stack := []int{}
	for index, v := range height {
		// 单调递减栈
		for 0 < len(stack) && height[stack[len(stack)-1]] < v { // 大于栈顶
			cur := stack[len(stack)-1] // 洼地
			stack = stack[:len(stack)-1]
			if 0 == len(stack) { // 没有 左挡板了  直接结束
				break
			}
			left := stack[len(stack)-1]             // 左挡板
			w := index - left - 1                   // 宽度
			h := min(height[left], v) - height[cur] // 高度
			res += w * h
		}
		stack = append(stack, index)
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```

### [统计矩形](https://leetcode-cn.com/problems/largest-rectangle-in-histogram)

```go
func largestRectangleArea(heights []int) int {
	n := len(heights)
	l, r := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		r[i] = n
	}
	stack := []int{} // 单调递增
	for i := 0; i < n; i++ {
		for 0 < len(stack) && heights[i] <= heights[stack[len(stack)-1]] {
			r[stack[len(stack)-1]] = i //  pop 时 计算右边界 ，只算当前到left的面积  相同的高 交给后面计算
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 { // 最小值 左边界为-1
			l[i] = -1
		} else {
			l[i] = stack[len(stack)-1] // push 时 左边界 为上一个
		}
		stack = append(stack, i)
	}
	res := 0
	for i := 0; i < n; i++ {
		res = max(res, (r[i]-l[i]-1)*heights[i])
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
```