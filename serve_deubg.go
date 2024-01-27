//go:build debug
// +build debug

package main

type DebugAppServer struct{}

func NewDebugAppServer() *DebugAppServer {
	return &DebugAppServer{}
}

func (s *DebugAppServer) Serve() error {
	return nil
}

func (s *DebugAppServer) Close() error {
	return nil
}

func (s *DebugAppServer) Addr() string {
	return "localhost:3000"
}
