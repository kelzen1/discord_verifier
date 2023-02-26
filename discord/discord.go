package discord

import (
	"Verifier/discord/interactions"
	discordModels "Verifier/models/discord"
	"fmt"
	"github.com/andersfylling/disgord"
)

func Init() {
	client := discordModels.GetClient()
	defer client.Gateway().StayConnectedUntilInterrupted()

	u, err := client.BotAuthorizeURL(disgord.PermissionUseSlashCommands, []string{
		"bot",
		"applications.commands",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(u)

	client.Gateway().BotReady(func() {
		go interactions.Setup(client)
	})

	client.Gateway().InteractionCreate(interactions.Handle)

}
