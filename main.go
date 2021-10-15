package main

import (
	"fmt"
)

type Test interface {
	A(int) int
}

type AFunc func(int) int

func A(n int) int {
	return n + 1
}

func b(a AFunc, n int) int {
	tmp := a(n)
	return tmp + 2
}

func main() {
	fmt.Println(b(A, 3))
}
