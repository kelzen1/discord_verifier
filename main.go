package main

import (
	"Verifier/database"
	"Verifier/discord"
	"Verifier/rest"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	go rest.Init()
	go database.Get()
	go discord.Init()

	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	<-s
}
