package main

import (
	"context"
	"log"
	"time"

	"github.com/bejaneps/rollee-assignment/test2/internal/service"
	"github.com/bejaneps/rollee-assignment/test2/internal/storage/inmemory"
	"github.com/bejaneps/rollee-assignment/test2/internal/transport"
	"github.com/bejaneps/rollee-assignment/test2/pkg/runtime"
	"github.com/caarlos0/env/v6"
)

const timeout = 5 * time.Second

type config struct {
	Port               string `env:"PORT" envDefault:"7171"`
	ServerReadTimeout  int64  `env:"SERVER_READ_TIMEOUT"`
	ServerWriteTimeout int64  `env:"SERVER_WRITE_TIMEOUT"`
}

func main() {
	cfg := &config{}
	err := env.Parse(cfg)
	if err != nil {
		log.Fatal(err)
	}

	storage := inmemory.New()

	mfwService := service.New(storage)

	var opts []transport.Option
	if cfg.ServerReadTimeout != 0 {
		opts = append(opts, transport.WithReadTimeout(cfg.ServerReadTimeout))
	}
	if cfg.ServerWriteTimeout != 0 {
		opts = append(opts, transport.WithWriteTimeout(cfg.ServerWriteTimeout))
	}

	mfwHandler := transport.New(
		mfwService,
		cfg.Port,
		opts...,
	)

	runtime.RunUntilSignal(
		func() error {
			return mfwHandler.Serve()
		},
		func(ctx context.Context) error {
			return mfwHandler.Close(ctx)
		},
		timeout,
	)
}
