package handlers

import "fmt"

func handleInit() (string, error) {
	fmt.Println("init")

	return "Done", nil
}
