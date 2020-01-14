package main

import (
	"testing"

	"github.com/mrojasb2000/building-microservices-with-go/rpc/client"
	"github.com/mrojasb2000/building-microservices-with-go/rpc/server"
)

func BenchmarkDial(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		c := client.CreateClient()
		c.Close()
	}
}

// go test -v -run="none" -bench=. -benchtime="5s"
func BenchmarkHelloWorldHandler(b *testing.B) {
	b.ResetTimer()

	c := client.CreateClient()

	for i := 0; i < b.N; i++ {
		client.PerformRequest(c)
	}

	c.Close()
}

func init() {
	go server.StartServer()
}
