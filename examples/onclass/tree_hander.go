package main

import "net/http"

type HandlerBaseOnTree struct {
	root *node
}

type node struct {
	path     string
	children []*node

	handler http.HandlerFunc
}

func main() {

}
