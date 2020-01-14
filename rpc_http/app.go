package main

import (
	"fmt"

	"github.com/mrojasb2000/building-microservices-with-go/rpc_http/client"
	"github.com/mrojasb2000/building-microservices-with-go/rpc_http/server"
)

func main() {
	server.StartServer()

	c := client.CreateClient()
	defer c.Close()

	reply := client.PerformRequest(c)

	fmt.Println(reply.Message)
}
