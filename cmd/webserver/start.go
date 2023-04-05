package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
)

func startServer(ctx context.Context) {
	e := initializeRouter()

	conf, err := newConfig()

	if err != nil {
		e.Logger.Fatal(err)
	}

	services, err := initializeServices(ctx, conf)

	if err != nil {
		e.Logger.Fatal(err)
	}

	registerRouters(e, services)

	svr := http.Server{
		Addr:    fmt.Sprintf(":%d", conf.Port),
		Handler: e,
		BaseContext: func(l net.Listener) context.Context {
			return ctx
		},
	}

	e.Logger.Fatal(e.StartServer(&svr))
}
