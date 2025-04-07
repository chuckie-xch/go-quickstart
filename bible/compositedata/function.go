package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	//funcDemo()
	//testAdd1()
	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

func square(n int) int {
	return n * n
}

// 形参引用类型，如指针,slice,map,function,channel会影响实参

// 函数之间是不可以比较的，也不能作为map的key

// 匿名函数被递归调用时要分成两步

func funcDemo() {
	var f func(int) int
	// f 为 nil 会引发 panic
	//f(3)
	f = square
	if f != nil {
		f(3)
	}

	fmt.Println(f(3))
	fmt.Printf("%T\n", f)

}

func add1(r rune) rune {
	return r + 1
}

func testAdd1() {
	fmt.Println(strings.Map(add1, "HAL-9000"))
	fmt.Println(strings.Map(add1, "VMS"))
}

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}
