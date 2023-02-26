package discordModels

import (
	"Verifier/utils"
	"github.com/andersfylling/disgord"
	"os"
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
			BotToken: os.Getenv("DISCORD_TOKEN"),
			Logger:   utils.Logger(),
		})
	})

	return client
}
