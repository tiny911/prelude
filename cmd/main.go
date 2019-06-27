package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"prelude/pkg/common/config"
	"prelude/pkg/common/server"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("[main] recover painc:%s.", r)
			exit(-1)
		}
	}()

	go func() {
		singals := make(chan os.Signal)
		signal.Notify(singals,
			syscall.SIGTERM,
			syscall.SIGINT,
			syscall.SIGKILL,
			syscall.SIGHUP,
			syscall.SIGQUIT,
		)
		<-singals
		exit(0)
	}()

	server.Run(config.Cfg.FrontPort)
}

func exit(status int) {
	server.Stop()
	log.Printf("[main] exit:%d.", status)
	os.Exit(status)
}
