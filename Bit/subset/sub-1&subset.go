package main

import "fmt"

func main() {
	subset := 0b00100101
	/** TODO
	00100100
	00100001
	00100000
	00000101
	00000100
	00000001
	00000000
	*/
	sub := subset
	for {
		if sub == 0 {
			break
		}
		sub = (sub - 1) & subset
		fmt.Printf("%08b\n", sub)
		if sub == subset {
			break
		}
	}
}
