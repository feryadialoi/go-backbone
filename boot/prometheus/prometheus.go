package prometheus

import (
	"fmt"
	"log"
)

type Prometheus struct {
	port int
}

func (p *Prometheus) SetPort(port int) {
	p.port = port
}

func (p *Prometheus) Run() error {
	log.Println(fmt.Sprintf("PROMETHEUS: run on port %v", p.port))
	return nil
}

func NewPrometheus() *Prometheus {
	return &Prometheus{}
}

func (p *Prometheus) Shutdown() {
	log.Println("PROMETHEUS: shutdown")
}
