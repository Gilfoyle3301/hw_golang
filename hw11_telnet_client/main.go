package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	timeout := flag.Duration("timeout", 10*time.Second, "connecction timeout")
	flag.Parse()
	if len(flag.Args()) != 2 {
		log.Fatalf("expected 2 arguments but got %d", len(flag.Args()))
	}
	host := flag.Arg(0)
	port := flag.Arg(1)
	newCliet := NewTelnetClient(net.JoinHostPort(host, port), *timeout, os.Stdin, os.Stdout)
	if err := newCliet.Connect(); err != nil {
		log.Fatalln(err)
	}
	ctxClient, cancel := signal.NotifyContext(context.Background(), syscall.SIGHUP, syscall.SIGINT)
	defer cancel()
	go func() {
		if err := newCliet.Send(); err != nil {
			log.Printf("Send bytes error: %v", err)
		}
		cancel()
	}()
	go func() {
		if err := newCliet.Receive(); err != nil {
			log.Printf("Receive bytes error: %v", err)
		}
		cancel()
	}()
	<-ctxClient.Done()
}
