package client

import (
	"fmt"
	"log"
	"net/rpc"

	"github.com/mrojasb2000/building-microservices-with-go/rpc_http/contract"
)

const port = 1234

// CreateClient create new instance of RPC client.
func CreateClient() *rpc.Client {
	client, err := rpc.DialHTTP("tpc", fmt.Sprintf("localhost:%v", port))
	if err != nil {
		log.Fatal("dialing:", err)
	}
	return client
}

// PerformRequest request compose
func PerformRequest(c *rpc.Client) contract.HelloWorldResponse {
	args := &contract.HelloWorldRequest{Name: "World"}
	var reply contract.HelloWorldResponse
	err := c.Call("HelloWorldHandler.HelloWorld", args, &reply)
	if err != nil {
		log.Fatal("error:", err)
	}
	return reply
}
