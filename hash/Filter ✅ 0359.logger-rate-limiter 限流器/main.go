package main

// https://leetcode-cn.com/problems/logger-rate-limiter

type Logger struct {
	m map[string]int
}

func Constructor() Logger {
	return Logger{m: make(map[string]int)}
}

func (l *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	if timestamp < l.m[message] {
		return false
	}
	l.m[message] = timestamp + 10
	return true
}
