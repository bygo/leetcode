package main

// https://leetcode-cn.com/problems/logger-rate-limiter

type Logger struct {
	m map[string]int
}

/** Initialize your data structure here. */
func Constructor() Logger {
	return Logger{m: make(map[string]int)}
}

/** Returns true if the message should be printed in the given timestamp, otherwise returns false.
  If this method returns false, the message will not be printed.
  The timestamp is in seconds granularity. */
func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	if timestamp < this.m[message] {
		return false
	}
	this.m[message] = timestamp + 10
	return true
}

/**
 * Your Logger object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.ShouldPrintMessage(timestamp,message);
 */
