package main

import "fmt"

func main() {
	//arrTest()
	testModifyArray()
}

func arrTest() {
	var a [3]int
	fmt.Println(a[0])

	for i, v := range a {
		fmt.Printf("%d %d \n", i, v)
	}

	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Println(r[1])
	fmt.Println(q[2])

	b := [...]int{99: -1}
	for _, v := range b {
		fmt.Println(v)
	}
}

func testModifyArray() {
	a := [5]int{1, 2, 3, 4, 5}
	modifyArray(a)
	fmt.Println(a[0])

	modifyArray2(&a)
	fmt.Println(a[0])
}

func modifyArray(arr [5]int) {
	arr[0] = 100
}

func modifyArray2(arr *[5]int) {
	arr[0] = 100
}
