package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type person struct {
	name   string
	age    int
	onWork bool
}

func main() {
	p1 := person{name: "tom", age: 100}
	p2 := person{name: "jack", age: 99}
	p3 := person{name: "sam", age: 20}
	people := map[string]*person{
		p1.name: &p1,
		p2.name: &p2,
		p3.name: &p3,
	}
	whoOnWork(people)
	if p3.onWork {
		fmt.Println("who is working?", p3.name)
	}
}

func whoOnWork(people map[string]*person) {
	for name, _ := range people {
		if people[name].age < 50 {
			//println(people[name].onWork)
			(people[name]).onWork = true
		}
	}
}
