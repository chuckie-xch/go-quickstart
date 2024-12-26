package main

import (
	"fmt"
	"net/http"
)

type Swimming interface {
	Swim()
}

type Duck interface {
	Swimming

	http.Handler

	MyMethod() string
}

type Base struct {
	Name string
}

type Concrete struct {
	Base
}

func (c *Concrete) sayHello() {

	fmt.Printf("I am base and my name is %s \n", c.Name)
	fmt.Printf("I am base and my name is: %s \n", c.Base.Name)

}
