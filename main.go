package main

import (
	"flag"

	"github.com/Surya-7890/gateway/server"
	"github.com/Surya-7890/gateway/server/api"
)

func main() {
	port := flag.String("port", ":7000", "the port on which the application runs") // port number as flag "go run main.go -port :8080"
	flag.Parse()
	server.Init()
	newServer := api.NewServer(*port)
	newServer.StartServer()
}
