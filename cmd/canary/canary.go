package main

import "github.com/getaddrinfo/canary/pkg/server"

func main() {
	// TODO: launch the control server as well
	server.NewServer(nil).Serve()
}
