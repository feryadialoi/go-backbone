package rest

import (
	"fmt"
	"log"
)

type Rest struct {
	port int
}

func NewRest() *Rest {
	return &Rest{}
}

func (r *Rest) SetPort(port int) {
	r.port = port
}

func (r *Rest) Run() error {
	log.Println(fmt.Sprintf("REST: run on port %v", r.port))
	return nil
}

func (r *Rest) Shutdown() {
	log.Println("REST: shutdown")
}
