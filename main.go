package main

import (
	"fmt"
	"log"

	"github.com/zserge/lorca"
)

func main() {
	ui, err := lorca.New("", "", 1080, 650, "--remote-allow-origins=*")
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()

	// Bind Go function to be available in JS. Go function may be long-running and
	// blocking - in JS it's represented with a Promise.
	ui.Bind("add", func(a, b int) int { return a + b })

	server := NewAppServer()
	err = server.Serve()
	if err != nil {
		log.Fatal(err)
	}
	defer server.Close()

	fmt.Println("Server started on", server.Addr())
	ui.Load(fmt.Sprintf("http://%s", server.Addr()))

	// Call JS function from Go. Functions may be asynchronous, i.e. return promises
	n := ui.Eval(`Math.random()`).Float()
	fmt.Println(n)

	// Call JS that calls Go and so on and so on...
	m := ui.Eval(`add(2, 3)`).Int()
	fmt.Println(m)

	// Wait for the browser window to be closed
	<-ui.Done()
}
