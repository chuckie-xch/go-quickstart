package main

import "fmt"

func main() {

}

func recoverDemo() {

}

func Parse(input string) (err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("internal error: %v", p)
		}
	}()
	return
}
