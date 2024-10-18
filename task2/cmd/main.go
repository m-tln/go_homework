package main

import (
	client "task2/internal/client"
	server "task2/internal/server"
)

func main() {
	server.StartServer()
	client.RunClient()
}
