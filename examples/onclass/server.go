package main

import "net/http"

type Server interface {
	Route(pattern string, handleFunc http.HandlerFunc)

	Start(address string) error
}

type SdkHttpServer struct {
	Name string
}

func (s *SdkHttpServer) Route(pattern string, handleFunc http.HandlerFunc) {
	http.HandleFunc(pattern, handleFunc)
}

func (s *SdkHttpServer) Start(address string) error {
	return http.ListenAndServe(address, nil)
}
