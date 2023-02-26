package discord

import (
	"github.com/andersfylling/disgord"
	"github.com/yoonaowo/discord_verifier/internal/discord/interactions"
	discordModels "github.com/yoonaowo/discord_verifier/internal/models/discord"
	"github.com/yoonaowo/discord_verifier/internal/utils"
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
