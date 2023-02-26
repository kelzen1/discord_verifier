package interactions

import (
	"github.com/andersfylling/disgord"

	"github.com/yoonaowo/discord_verifier/internal/models"
	discordModels "github.com/yoonaowo/discord_verifier/internal/models/discord"
	utils2 "github.com/yoonaowo/discord_verifier/internal/utils"

	"context"
	"os"
	"sync"
)

var interactions = &[]*models.Interaction{}
var mapInteractions = make(map[string]*models.Interaction)
var mutex sync.Mutex

func AddInteraction(interaction *models.Interaction) {

	mutex.Lock()
	defer mutex.Unlock()

	*interactions = append(*interactions, interaction)
	mapInteractions[interaction.CommandDefinition.Name] = interaction
}

func GetInteraction(name string) (interactionData *models.Interaction, err error) {

	mutex.Lock()
	defer mutex.Unlock()

	err = utils2.ErrInteractionNotFound

	interaction, ok := mapInteractions[name]

	if !ok {
		return
	}

	return interaction, nil
}

var failedResponse = &disgord.CreateInteractionResponse{
	Type: disgord.InteractionCallbackChannelMessageWithSource,
	Data: &disgord.CreateInteractionResponseData{
		Embeds: []*disgord.Embed{
			{
				Title:       "Error",
				Description: "Aww, something went wrong!~",
			},
		},
	},
}

func handle(session disgord.Session, interactionCreate *disgord.InteractionCreate) {
	interaction, err := GetInteraction(interactionCreate.Data.Name)

	if err != nil {
		_ = discordModels.GetClient().SendInteractionResponse(context.Background(), interactionCreate, failedResponse)
		return
	}

	interaction.Callback(session, interactionCreate)
}

func Handle(session disgord.Session, interactionCreate *disgord.InteractionCreate) {
	go handle(session, interactionCreate) // memes
}

func Setup(client *disgord.Client) {
	AddInteraction(verifyStruct)

	GuildID, err := disgord.GetSnowflake(os.Getenv("DISCORD_GUILD"))

	if err != nil {
		utils2.Logger().Panicln("get snowflake interactions ->", err)
		return
	}

	for _, interaction := range *interactions {
		if err := client.ApplicationCommand(0).Guild(GuildID).Create(interaction.CommandDefinition); err != nil {
			utils2.Logger().Fatalf(err.Error())
		}
	}

}
