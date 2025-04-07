package main

import (
	"fmt"
)

func main() {
	pointer()
	pointNew()
}

func statement() {
	var s string
	fmt.Println(s)

	var i, j, k int
	fmt.Println(i, j, k)
	var b, f, s1 = true, 2.3, "four"
	fmt.Println(b, f, s1)

	t := 0.0
	fmt.Println(t)
}

func pointer() {
	x := 1
	p := &x
	fmt.Println(*p)
	*p = 2
	fmt.Println(x)

}

func pointNew() {
	p := new(int)
	fmt.Println(*p)

	*p = 2
	fmt.Println(*p)
}
