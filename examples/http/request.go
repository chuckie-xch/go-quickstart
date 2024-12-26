package main

import (
	"fmt"
	"net/http"
)

func header(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "header is %v\n", r.Header)

}

func main() {
	http.HandleFunc("header", header)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}

}
