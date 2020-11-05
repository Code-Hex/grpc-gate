package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	gate "github.com/Code-Hex/grpc-gate"
	"github.com/pkg/errors"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Println(err)
	}
}

func run(ctx context.Context) error {
	ln, err := net.Listen("tcp", "127.0.0.1:4000")
	if err != nil {
		return errors.WithStack(err)
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT)

	srv := gate.NewHandler()
	go func() {
		log.Println("serve =>", ln.Addr().String())
		if err := srv.Serve(ln); err != nil {
			log.Println(err)
			return
		}
	}()
	select {
	case <-c:
		log.Println("shutdown...")
		srv.Stop()
	}
	return nil
}
