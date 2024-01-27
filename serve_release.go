//go:generate go run -tags generate gen.go
//go:build !debug
// +build !debug

package main

import (
	"net"
	"net/http"
)

type ProductionAppServer struct {
	ln net.Listener
}

func NewProductionAppServer() *ProductionAppServer {
	return &ProductionAppServer{}
}

func (s *ProductionAppServer) Serve() error {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return err
	}
	s.ln = ln
	server := http.NewServeMux()

	server.Handle("/", http.StripPrefix("/", http.FileServer(FS)))
	go http.Serve(ln, server)
	return nil
}

func (s *ProductionAppServer) Close() error {
	if s.ln != nil {
		return s.ln.Close()
	}
	return nil
}

func (s *ProductionAppServer) Addr() string {
	return s.ln.Addr().String()
}
