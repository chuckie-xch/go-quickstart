package main

import "time"

type Employee struct {
	Id        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerId int
}

var dilbert Employee

func demo() {
	dilbert.Id = 1
	position := &dilbert.Position
	*position = "Senior"

	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position = " (proactive team player)"

}

type tree struct {
	value       int
	left, right *tree
}

func Sort(values []int) {
	var root *tree
	for _, val := range values {
		root = add(root, val)
	}
	appendValues(values[:0], root)
}

func appendValues(values []int, root *tree) []int {
	if root != nil {
		values = appendValues(values, root.left)
		values = append(values, root.value)
		values = appendValues(values, root.right)
	}
	return values
}

func add(root *tree, value int) *tree {
	if root == nil {
		root = new(tree)
		root.value = value
		return root
	}
	if root.value > value {
		root.left = add(root.left, value)
	} else {
		root.right = add(root.right, value)
	}
	return root
}

func setIfAbsent(key string) {
	seen := make(map[string]struct{})
	// ...
	if _, ok := seen[key]; !ok {
		seen[key] = struct{}{}
	}
}
