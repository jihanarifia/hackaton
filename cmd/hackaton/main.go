package main

import (
	"flag"
	"fmt"
	"hackaton/pkg/config"
	"hackaton/pkg/server"
	"hackaton/pkg/service"
	"os"
	"os/signal"
	"syscall"

	"github.com/caarlos0/env"
	"github.com/pkg/errors"
)

func main() {
	conf := config.Config{}
	flag.Usage = func() {
		flag.CommandLine.SetOutput(os.Stdout)
		for _, val := range conf.HelpDocs() {
			fmt.Println(val)
		}
		fmt.Println("")
		flag.PrintDefaults()
	}
	flag.Parse()

	err := env.Parse(&conf)
	if err != nil {
		fmt.Printf("%+v\n", errors.WithStack(err))
		return
	}

	if err != nil {
		fmt.Printf("%+v\n", errors.WithStack(err))
		return
	}

	restfulService := service.New(server.ServiceName, conf)

	restfulServer, err := server.New(&conf, restfulService)
	if err != nil {
		fmt.Printf("%+v\n", errors.WithStack(err))
		return
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sig
		restfulServer.Stop()
	}()

	restfulServer.Serve()
}
