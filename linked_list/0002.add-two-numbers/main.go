package main

/**
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l := l1
	var carry int
	for {
		l1.Val += l2.Val + carry
		carry = l1.Val / 10
		l1.Val = l1.Val % 10
		if l2.Next == nil {
			if carry == 0 {
				break
			}
			l2.Next = &ListNode{}
		}
		if l1.Next == nil {
			if carry == 0 {
				l1.Next = l2.Next
				break
			}
			l1.Next = &ListNode{}
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l
}



/**
思路：
1.避免创建新链表，以 l1 or l2 进行更改
2.其一链表Next指向为空时，开始尝试链表合并
 */
