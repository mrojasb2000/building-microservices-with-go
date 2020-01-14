package server

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/mrojasb2000/building-microservices-with-go/rpc_http/contract"
)

const port = 1234

// HTTPConn type
type HTTPConn struct {
	in  io.Reader
	out io.Writer
}

// Read
func (c *HTTPConn) Read(p []byte) (n int, err error)  { return c.in.Read(p) }
func (c *HTTPConn) Write(d []byte) (n int, err error) { return c.out.Write(d) }

// Close connection
func (c *HTTPConn) Close() error { return nil }

// HelloWorldHandler handler type
type HelloWorldHandler struct{}

// HelloWorld method for handler
func (h *HelloWorldHandler) HelloWorld(args *contract.HelloWorldRequest, reply *contract.HelloWorldResponse) error {
	reply.Message = "Hello " + args.Name
	return nil
}

// StartServer initialize RPC server
func StartServer() {
	helloWorld := &HelloWorldHandler{}
	rpc.Register(helloWorld)

	l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to listen on given port %s", err))
	}

	log.Printf("Server starting on port %v\n", port)

	http.Serve(l, http.HandlerFunc(httpHandler))
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	serverCodec := jsonrpc.NewServerCodec(&HTTPConn{in: r.Body, out: w})
	err := rpc.ServeRequest(serverCodec)
	if err != nil {
		log.Printf("Error while serving JSON request: %v", err)
		http.Error(w, "Error while serving JSON request, details have been logged.", 500)
		return
	}
}
