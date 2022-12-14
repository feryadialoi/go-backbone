package main

import (
	"context"
	"github.com/feryadialoi/demo/boot"
	"github.com/feryadialoi/demo/boot/grpc"
	"github.com/feryadialoi/demo/boot/prometheus"
	"github.com/feryadialoi/demo/boot/rest"
	"github.com/feryadialoi/demo/config"
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
