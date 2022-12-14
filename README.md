# Go Backbone

Example of golang framework with graceful

### main.go
```go
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
```

### Log example
```shell
2022/12/14 20:59:35 BOOT: run
2022/12/14 20:59:35 PROMETHEUS: run on port 9800
2022/12/14 20:59:35 GRPC: run on port 9090
2022/12/14 20:59:35 REST: run on port 8080
2022/12/14 20:59:45 shutting down...
2022/12/14 20:59:45 PROMETHEUS: shutdown
2022/12/14 20:59:45 REST: shutdown      
2022/12/14 20:59:47 GRPC: shutdown
2022/12/14 20:59:47 application shutdown gracefully

Process finished with the exit code 0
```