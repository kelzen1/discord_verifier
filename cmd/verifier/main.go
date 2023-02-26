package main

import (
	"github.com/yoonaowo/discord_verifier/internal/database"
	"github.com/yoonaowo/discord_verifier/internal/discord"
	"github.com/yoonaowo/discord_verifier/internal/rest"
	"github.com/yoonaowo/discord_verifier/internal/utils"
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
