package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/tiny911/prelude/internal/common/config"
	"github.com/tiny911/prelude/internal/common/server"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("[main] recover painc:%s.", r)
			exit(-1)
		}
	}()

	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals,
			syscall.SIGTERM,
			syscall.SIGINT,
			// syscall.SIGKILL,
			syscall.SIGHUP,
			syscall.SIGQUIT,
		)
		<-signals
		exit(0)
	}()

	server.Run(config.Cfg.FrontPort)
}

func exit(status int) {
	server.Stop()
	log.Printf("[main] exit:%d.", status)
	os.Exit(status)
}
