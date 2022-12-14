package boot

import (
	"context"
	"github.com/feryadialoi/demo/boot/grpc"
	"github.com/feryadialoi/demo/boot/prometheus"
	"github.com/feryadialoi/demo/boot/rest"
	"log"
	"sync"
)

type Boot struct {
	config *Config
	apps   map[string]RunShutdowner
}

func NewBoot(applyConfigs ...ApplyConfig) *Boot {
	c := DefaultConfig()

	for _, applyConfig := range applyConfigs {
		applyConfig(c)
	}

	return &Boot{
		config: c,
		apps:   make(map[string]RunShutdowner),
	}
}

func (b *Boot) Run(ctx context.Context) error {
	log.Println("BOOT: run")
	for _, r := range b.apps {
		if err := r.Run(); err != nil {
			return err
		}
	}

	return nil
}

func (b *Boot) RegisterGRPC(grpc *grpc.GRPC) {
	grpc.SetPort(b.config.grpcPort)
	b.apps["grpc"] = grpc
}

func (b *Boot) RegisterRest(rest *rest.Rest) {
	rest.SetPort(b.config.restPort)
	b.apps["rest"] = rest
}

func (b *Boot) RegisterPrometheus(prometheus *prometheus.Prometheus) {
	prometheus.SetPort(b.config.prometheusPort)
	b.apps["prometheus"] = prometheus
}

func (b *Boot) Shutdown() {
	wg := sync.WaitGroup{}
	for _, s := range b.apps {
		wg.Add(1)
		go func(s Shutdowner) {
			s.Shutdown()
			wg.Done()
		}(s)
	}
	wg.Wait()
}
