package main

import "errors"

func main() {
	var err error = &MyError{}

	println(err.Error())

	err = errors.New("我自己的Error")

	println(err.Error())

}

type MyError struct {
}

func (m *MyError) Error() string {
	return "Hello, it's my error"

}
