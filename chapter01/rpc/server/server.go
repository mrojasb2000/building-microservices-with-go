package server

import (
	"fmt"
	"log"
	"net"
	"net/rpc"

	"github.com/mrojasb2000/building-microservices-with-go/rpc/contract"
)

const port = 1234

func main() {
	log.Printf("Server staring on port: %v\n", port)
	StartServer()
}

// StartServer initialize RPC server.
func StartServer() {
	helloWorld := &HelloWorldHandler{}
	rpc.Register(helloWorld)

	l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to listen on given port: %s", err))
	}

	for {
		conn, _ := l.Accept()
		go rpc.ServeConn(conn)
	}
}

// HelloWorldHandler is handler
type HelloWorldHandler struct{}

// HelloWorld function
func (h *HelloWorldHandler) HelloWorld(args *contract.HelloWorldRequest,
	reply *contract.HelloWorldResponse) error {
	reply.Message = "Hello " + args.Name
	return nil
}
