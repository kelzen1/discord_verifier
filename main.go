package main

import (
	"Verifier/database"
	"Verifier/discord"
	"Verifier/rest"
	"Verifier/utils"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	defer utils.ShutdownLogger()

	go rest.Init()
	go database.Get()
	go discord.Init()

	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	<-s
}
