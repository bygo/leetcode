package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//Title: Design Phone Directory
//Link: https://leetcode-cn.com/problems/design-phone-directory

type PhoneDirectory struct {
	start, max int
	Used       map[int]struct{}
	Unused     []int
}

/** Initialize your data structure here
  @param maxNumbers - The maximum numbers that can be stored in the phone directory. */
func Constructor(maxNumbers int) PhoneDirectory {
	return PhoneDirectory{
		max:  maxNumbers,
		Used: map[int]struct{}{},
	}
}

/** Provide a number which is not assigned to anyone.
  @return - Return an available number. Return -1 if none is available. */
func (this *PhoneDirectory) Get() int {
	var num = -1
	if this.start < this.max {
		num = this.start
		this.start++
		this.Used[num] = struct{}{}
	} else if 0 < len(this.Unused) {
		num = this.Unused[0]
		this.Unused = this.Unused[1:]
		this.Used[num] = struct{}{}
	}

	return num
}

/** Check if a number is available or not. */
func (this *PhoneDirectory) Check(number int) bool {
	_, ok := this.Used[number]
	return !ok
}

/** Recycle or release a number. */
func (this *PhoneDirectory) Release(number int) {
	if !this.Check(number) {
		this.Unused = append(this.Unused, number)
		delete(this.Used, number)
	}

	if 0 == len(this.Used) {
		this.Unused = []int{}
		this.start = 0
	}
}

/**
 * Your PhoneDirectory object will be instantiated and called as such:
 * obj := Constructor(maxNumbers);
 * param_1 := obj.Get();
 * param_2 := obj.Check(number);
 * obj.Release(number);
 */
