package server

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/mrojasb2000/building-microservices-with-go/rpc_http/contract"
)

const port = 1234

// HelloWorldHandler handler
type HelloWorldHandler struct{}

// HelloWorld method of HelloWorldHandler
func (h HelloWorldHandler) HelloWorld(args *contract.HelloWorldRequest, reply *contract.HelloWorldResponse) error {
	reply.Message = "Hello " + args.Name
	return nil
}

// StartServer initialize RPC over HTTP server
func StartServer() {
	helloWorld := &HelloWorldHandler{}
	rpc.Register(helloWorld)
	rpc.HandleHTTP()

	l, err := net.Listen("tpc", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to listen on given port %s", err))
	}

	log.Printf("Server starting on port %v\n", port)

	http.Serve(l, nil)

}
