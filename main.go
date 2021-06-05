package main

import "fmt"

type interface1 interface {
	me1()
	me2()
}

type A struct {
	a int
}

func (a A) me1() int {
	return 1
}

func main() {
	a := &A{}
	fmt.Println(a.me1())
}
