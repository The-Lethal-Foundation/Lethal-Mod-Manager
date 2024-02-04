//go:build !debug

//go:generate go run -tags generate gen.go
package main

import (
	"fmt"
	"net"
	"net/http"
	"path/filepath"

	"github.com/The-Lethal-Foundation/lethal-core/filesystem"
)

type AppServer struct {
	ln net.Listener
}

func NewAppServer() *AppServer {
	fmt.Println("RELEASE MODE")
	return &AppServer{}
}

func (s *AppServer) Serve() error {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return err
	}
	s.ln = ln
	server := http.NewServeMux()

	profilesPath := filepath.Join(filesystem.GetDefaultPath(), "LethalCompany", "Profiles")

	server.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir(profilesPath))))
	server.Handle("/", http.StripPrefix("/", http.FileServer(FS)))

	go http.Serve(ln, server)
	return nil
}

func (s *AppServer) Close() error {
	if s.ln != nil {
		return s.ln.Close()
	}
	return nil
}

func (s *AppServer) Addr() string {
	return s.ln.Addr().String()
}
