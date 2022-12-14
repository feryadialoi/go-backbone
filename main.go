package main

import (
	"context"
	"github.com/feryadialoi/go-backbone/boot"
	"github.com/feryadialoi/go-backbone/boot/grpc"
	"github.com/feryadialoi/go-backbone/boot/prometheus"
	"github.com/feryadialoi/go-backbone/boot/rest"
	"github.com/feryadialoi/go-backbone/config"
)

func main() {
	c := config.GetConfig()

	b := boot.NewBoot(
		boot.RestPort(c.RestPort()),
		boot.GRPCPort(c.GRPCPort()),
		boot.PrometheusPort(c.PrometheusPort()),
	)

	grpcSrv := grpc.NewGRPC()
	restSrv := rest.NewRest()
	prometheusSvc := prometheus.NewPrometheus()

	b.RegisterGRPC(grpcSrv)
	b.RegisterRest(restSrv)
	b.RegisterPrometheus(prometheusSvc)

	boot.RunGracefully(context.Background(), b, c.Timeout())
}
