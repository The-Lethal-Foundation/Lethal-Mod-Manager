package main

import (
	"fmt"
	"log"

	"github.com/KonstantinBelenko/lethal-mod-manager/internal/handlers"
	"github.com/zserge/lorca"
)

func main() {
	ui, err := lorca.New("", "", 1080, 650, "--remote-allow-origins=*")
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()

	server := NewAppServer()
	err = server.Serve()
	if err != nil {
		log.Fatal(err)
	}
	defer server.Close()

	fmt.Println("Server started on", server.Addr())
	ui.Load(fmt.Sprintf("http://%s", server.Addr()))

	// Initialize handlers
	handlers.SetupHandlers(ui)

	// Wait for the browser window to be closed
	<-ui.Done()
}
