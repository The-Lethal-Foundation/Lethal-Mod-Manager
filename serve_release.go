//go:build !debug

//go:generate go run -tags generate gen.go
package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/KonstantinBelenko/lethal-mod-manager/pkg/lcfs/util"
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

	profilesPath, err := util.GetProfilesPath()
	if err != nil {
		log.Fatal(err)
	}

	server.Handle("/", http.StripPrefix("/", http.FileServer(FS)))
	server.Handle("/images", http.StripPrefix("/images", http.FileServer(http.Dir(profilesPath))))

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
