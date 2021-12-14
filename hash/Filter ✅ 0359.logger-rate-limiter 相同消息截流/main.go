package main

// https://leetcode-cn.com/problems/logger-rate-limiter

// ❓相同消息截流

type Logger struct {
	msgMpTime map[string]int
}

func Constructor() Logger {
	return Logger{msgMpTime: make(map[string]int)}
}

func (l *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	// 还在截流期
	if timestamp < l.msgMpTime[message] {
		return false
	}
	// 10秒
	l.msgMpTime[message] = timestamp + 10
	return true
}
