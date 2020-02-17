package main

import (
	"testing"
)

var inputs = map[string]int{
	"   -41":       -41,
	" 123123 kkkk": 123123,
	" +12313":      12313,
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for k, _ := range inputs {
			myAtoi(k)
		}
	}
}

func Test(t *testing.T) {
	for k, v := range inputs {
		if myAtoi(k) != v {
			t.Error("Error")
		}
	}

}
