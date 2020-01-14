package main

import (
	"fmt"

	"github.com/mrojasb2000/building-microservices-with-go/rpc_http/client"
	"github.com/mrojasb2000/building-microservices-with-go/rpc_http/server"
)

func main() {
	fmt.Println("A0")
	server.StartServer()
	fmt.Println("A1")
	c := client.CreateClient()
	defer c.Close()
	fmt.Println("A2")

	reply := client.PerformRequest(c)
	fmt.Println("A3")
	fmt.Println(reply.Message)
}
