//go:build debug

package main

import "fmt"

type AppServer struct{}

func NewAppServer() *AppServer {
	fmt.Println("DEBUG MODE")
	return &AppServer{}
}

func (s *AppServer) Serve() error {
	return nil
}

func (s *AppServer) Close() error {
	return nil
}

func (s *AppServer) Addr() string {
	return "localhost:3000"
}
