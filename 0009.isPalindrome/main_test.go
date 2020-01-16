package main

import (
	"testing"
)

var inputs = map[int]bool{
	123:   false,
	12321: true,
	-123:  false,
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for k, _ := range inputs {
			isPalindrome(k)
		}
	}
}

func Test(t *testing.T) {
	for k, v := range inputs {
		if isPalindrome(k) != v {
			t.Error("Error")
		}
	}

}
