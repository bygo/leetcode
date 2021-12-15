# 前驱指针 zero

### [加一](https://leetcode-cn.com/problems/plus-one-linked-list)

```go
func plusOne(head *ListNode) *ListNode {
	zero := &ListNode{Next: head}
	cur := zero
	// 找非9
	for head != nil {
		if head.Val != 9 {
			cur = head
		}
		head = head.Next
	}

	cur.Val += 1

	// cur后 所有 9 转 0
	for cur.Next != nil {
		cur.Next.Val = 0
		cur = cur.Next
	}

	// 是否需要进位
	if zero.Val == 0 {
		return zero.Next
	} else {
		return zero
	}
}
```

### [删除重复全部元素](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii)

```go
func deleteDuplicates(head *ListNode) *ListNode {
	zero := &ListNode{Next: head}
	cur := zero
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val { // 出现相同
			v := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == v { // 所有相同的都删除
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}
	return zero.Next
}
```

# 快慢指针 slow fast

### [环的起点](https://leetcode-cn.com/problems/linked-list-cycle-ii)

```go
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast { // 相遇时 在环内
			for slow != head { // 如果以环起点，相遇时必定也是环起点
				slow = slow.Next // 因为 slow 入环  fast 相对于 slow 多走了 head->环起点
				head = head.Next // 所以 相遇时 slow head->环起点 = slow->环起点
			}
			return slow
		}
	}
	return nil
}

```

### [中间节点](https://leetcode-cn.com/problems/middle-of-the-linked-list)

```go
func middleNode(head *ListNode) *ListNode {
	var fast, slow = head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
```

# 哨兵指针 prev

### [两两交换节点](https://leetcode-cn.com/problems/swap-nodes-in-pairs)

```go
func swapPairs(head *ListNode) *ListNode {
	zero := &ListNode{Next: head}
	prev := zero
	for head != nil && head.Next != nil {
		l := head
		r := head.Next

		prev.Next = r
		l.Next = r.Next
		r.Next = l

		prev = l
		head = prev.Next
	}
	return zero.Next
}
```

### [合并两个有序链表](https://leetcode-cn.com/problems/merge-two-sorted-lists)

```go
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	zero := &ListNode{}
	prev := zero
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			prev.Next = l1
			l1 = l1.Next
		} else {
			prev.Next = l2
			l2 = l2.Next
		}
		prev = prev.Next
	}

	if l1 == nil {
		prev.Next = l2
	} else if l2 == nil {
		prev.Next = l1
	}
	return zero.Next
}
```