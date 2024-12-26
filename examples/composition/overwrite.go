package main

import "fmt"

func main() {
	son := Son{
		Parent{},
	}
	son.SayHello()
}

type Parent struct {
}

func (p Parent) SayHello() {
	fmt.Println("I am " + p.Name())
}

func (p Parent) Name() string {
	return "Parent"
}

type Son struct {
	Parent
}

func Name() string {
	return "Son"
}
