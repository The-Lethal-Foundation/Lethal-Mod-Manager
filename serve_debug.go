// serve_debug.go
//go:build debug

package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/KonstantinBelenko/lethal-mod-manager/pkg/lcfs/util"
)

type AppServer struct{}

func NewAppServer() *AppServer {
	fmt.Println("DEBUG MODE")
	return &AppServer{}
}

func (s *AppServer) Serve() error {
	server := http.NewServeMux()

	// Replicate the file serving logic for images
	profilesPath, err := util.GetProfilesPath()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n\n---------------------")
	fmt.Println("Serving images from", profilesPath)
	fmt.Println("---------------------\n\n")

	server.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir(profilesPath))))
	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.URL.Path, "/images/") {
			http.Redirect(w, r, "http://localhost:3000", http.StatusMovedPermanently)
			return
		}
		http.NotFound(w, r)
	})

	// Start the HTTP server on a predetermined debug port
	go func() {
		if err := http.ListenAndServe("localhost:3001", server); err != nil {
			log.Fatal(err)
		}
	}()
	return nil
}

func (s *AppServer) Close() error {
	return nil
}

func (s *AppServer) Addr() string {
	return "localhost:3001"
}
