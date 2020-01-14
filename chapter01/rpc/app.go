package main

import (
	"fmt"

	"github.com/mrojasb2000/building-microservices-with-go/rpc/client"
	"github.com/mrojasb2000/building-microservices-with-go/rpc/server"
)

func main() {
	go server.StartServer()

	c := client.CreateClient()
	defer c.Close()

	reply := client.PerformRequest(c)
	fmt.Println(reply.Message)

}
