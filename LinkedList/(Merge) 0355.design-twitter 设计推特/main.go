package main

// https://leetcode.cn/problems/design-twitter

type ListNode struct {
	TweetId int
	Time    int
	Next    *ListNode
}

type Twitter struct {
	Msg      map[int]*ListNode
	Follower map[int]map[int]struct{}
	Time     int
}

func Constructor() Twitter {
	return Twitter{
		Msg:      map[int]*ListNode{},
		Follower: map[int]map[int]struct{}{},
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	n := &ListNode{TweetId: tweetId, Time: this.Time}
	this.Time++
	n.Next = this.Msg[userId]
	this.Msg[userId] = n
	cnt := 10
	for n.Next != nil && 0 < cnt {
		cnt--
		n = n.Next
	}
	n.Next = nil
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	ls := []*ListNode{this.Msg[userId]}
	for id := range this.Follower[userId] {
		ls = append(ls, this.Msg[id])
	}
	l := mergeKLists(ls)
	var res []int
	cnt := 0
	for l != nil && cnt < 10 {
		res = append(res, l.TweetId)
		l = l.Next
		cnt++
	}
	return res
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if followerId == followeeId {
		return
	}
	if this.Follower[followerId] == nil {
		this.Follower[followerId] = map[int]struct{}{}
	}
	this.Follower[followerId][followeeId] = struct{}{}
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if followerId == followeeId {
		return
	}
	if this.Follower[followerId] == nil {
		this.Follower[followerId] = map[int]struct{}{}
	}
	delete(this.Follower[followerId], followeeId)
}

func mergeKLists(lists []*ListNode) *ListNode {
	l := len(lists)
	if 0 == l {
		return nil
	}

	for 1 < l {
		for i := 0; i < l/2; i++ {
			lists[i] = mergeTwoLists(lists[i], lists[l-i-1])
		}
		l = (l + 1) / 2
	}
	return lists[0]
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	prev := &ListNode{}
	res := prev
	cnt := 0
	for cnt < 10 && l1 != nil && l2 != nil {
		if l2.Time < l1.Time {
			prev.Next = &ListNode{
				TweetId: l1.TweetId,
				Time:    l1.Time,
			}
			l1 = l1.Next
		} else {
			prev.Next = &ListNode{
				TweetId: l2.TweetId,
				Time:    l2.Time,
			}
			l2 = l2.Next
		}
		prev = prev.Next
		cnt++
	}

	if l2 != nil {
		l1 = l2
	}

	for cnt < 10 && l1 != nil {
		prev.Next = &ListNode{
			TweetId: l1.TweetId,
			Time:    l1.Time,
		}
		prev = prev.Next
		l1 = l1.Next
		cnt++
	}
	return res.Next
}
