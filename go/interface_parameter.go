package main

import (
	"fmt"
)

type I interface {
	m()
} 

type T struct {}

func (t T) m() {
	fmt.Println("m()")
}

func myfunc(i I) {
	fmt.Printf("%#v\n", i)
	i.m()
}

func main() {
	t := T{}
	myfunc(t)
	myfunc(&t)
}
