package discordModels

import (
	"github.com/andersfylling/disgord"

	"github.com/yoonaowo/discord_verifier/internal/utils"

	"sync"
)

var (
	once   sync.Once
	client *disgord.Client
)

// this is singleton
func GetClient() *disgord.Client {
	once.Do(func() {
		client = disgord.New(disgord.Config{
			BotToken: utils.FlagDiscordToken,
			Logger:   utils.Logger(),
		})
	})

	return client
}
