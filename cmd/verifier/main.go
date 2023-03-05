package main

import (
	"github.com/yoonaowo/discord_verifier/internal/database"
	"github.com/yoonaowo/discord_verifier/internal/discord"
	"github.com/yoonaowo/discord_verifier/internal/rest"
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
