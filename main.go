package main

import "github.com/fmcarrero/contacts-api/src"

func main() {
	dependencies := src.Build()
	defer dependencies.Close()
	server := src.NewServer(dependencies)

	server.Routes()
	server.Start()

}
