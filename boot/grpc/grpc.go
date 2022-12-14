package grpc

import (
	"fmt"
	"log"
	"time"
)

type GRPC struct {
	port int
}

func NewGRPC() *GRPC {
	return &GRPC{}
}

func (g *GRPC) Run() error {
	log.Println(fmt.Sprintf("GRPC: run on port %v", g.port))
	return fmt.Errorf("GRPC: err")
}

func (g *GRPC) SetPort(port int) {
	g.port = port
}

func (g *GRPC) Shutdown() {
	time.Sleep(time.Second * 2)
	log.Println(fmt.Sprintf("GRPC: shutdown"))
}
