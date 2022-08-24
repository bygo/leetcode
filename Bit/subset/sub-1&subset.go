package main

import "fmt"

func main() {
	subset := 0b00100101
	sub := subset
	for {
		if sub == 0 {
			break
		}
		sub = (sub - 1) & subset
		fmt.Printf("%08b\n", sub)

		//if sub == subset {
		//	break
		//}
	}
}
