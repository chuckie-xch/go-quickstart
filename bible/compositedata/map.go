package main

import "fmt"

func main() {
	example()

}

func example() {
	ages := make(map[string]int)
	ages["alice"] = 31
	ages["join"] = 32

	delete(ages, "alice")

	age, ok := ages["bob"]
	if !ok {
		fmt.Printf("%v's age is not found", "bob")
	} else {
		fmt.Printf("bob's age is %d", age)
	}

	// 不存在时返回零值
	ages["bob"] = ages["bob"] + 1
}

func mapEquals(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}

	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}

	return true
}
