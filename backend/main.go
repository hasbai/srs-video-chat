package main

import (
	"github.com/hasbai/srs-video-chat/handler"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	httpAddr := ":8000"
	go handler.Serve(httpAddr)

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	s := <-c
	log.Println("received", s)
	exit()
}

func exit() {
	log.Printf("exiting...")
	os.Exit(0)
}
