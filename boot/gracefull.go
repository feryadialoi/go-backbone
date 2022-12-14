package boot

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func RunGracefully(ctx context.Context, boot *Boot, timeout time.Duration) {
	wait := make(chan struct{})
	go func() {
		if err := boot.Run(ctx); err != nil {
			wait <- struct{}{}
			log.Fatal(err)
		}
	}()

	go func() {
		s := make(chan os.Signal, 1)
		signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
		<-s

		log.Println("shutting down...")

		timeoutFunc := time.AfterFunc(timeout, func() {
			log.Printf("timeout %d ms has been elapsed, force exit\n", timeout.Milliseconds())
			os.Exit(0)
		})
		defer timeoutFunc.Stop()

		boot.Shutdown()

		wait <- struct{}{}
	}()

	<-wait

	log.Println("application shutdown gracefully")
}
