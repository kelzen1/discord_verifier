package discordModels

import (
	"github.com/andersfylling/disgord"
	"github.com/sirupsen/logrus"
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
			Logger:   logrus.New(),
		})
	})

	return client
}
