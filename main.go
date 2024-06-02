package main

import "github.com/fmcarrero/contacts-api/src"

func main() {
	dependencies := src.Build()
	defer dependencies.CloseDatabase()
	server := src.NewServer(dependencies)
	server.Middlewares(src.WithRecover(), src.WithRequestID(), src.WithErrorHandler())
	server.Routes()
	server.Start()
}
