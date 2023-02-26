package discord

import (
	"Verifier/discord/interactions"
	discordModels "Verifier/models/discord"
	"Verifier/utils"
	"github.com/andersfylling/disgord"
)

func Init() {
	client := discordModels.GetClient()
	defer client.Gateway().StayConnectedUntilInterrupted()

	userLink, err := client.BotAuthorizeURL(disgord.PermissionUseSlashCommands, []string{
		"bot",
		"applications.commands",
	})
	if err != nil {
		panic(err)
	}
	utils.Logger().Println("Invite link:", userLink)

	client.Gateway().BotReady(func() {
		go interactions.Setup(client)
	})

	client.Gateway().InteractionCreate(interactions.Handle)

}
