package main

import "fmt"

func main() {

	i := 13
	a := func() {
		fmt.Println("i is %d \n", i)
	}

	a()

	fmt.Println(ReturnClosure("Tom")())

	Delay()

}

func ReturnClosure(name string) func() string {
	return func() string {
		return "Hello, " + name
	}
}

func Delay() {
	fns := make([]func(), 0, 10)
	for i := 0; i < 10; i++ {
		fns = append(fns, func() {
			fmt.Printf("hello, this is : %d \n", i)
		})
	}

	for _, fn := range fns {
		fn()
	}
}
