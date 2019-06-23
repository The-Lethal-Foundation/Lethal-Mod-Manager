//+build debug

package main

func serve() (string, error) {
	// Assume react dev server runs on default address
	return "http://localhost:3000", nil
}
