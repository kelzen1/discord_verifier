package interactions

import (
	"Verifier/database"
	"Verifier/models"
	discordModels "Verifier/models/discord"
	"Verifier/utils"
	"context"
	"database/sql"
	"github.com/andersfylling/disgord"
)

var (
	verifyDefinition = &disgord.CreateApplicationCommand{
		Name:        "verify",
		Description: "Get User Role",
		Options: []*disgord.ApplicationCommandOption{
			{
				Required:    true,
				Name:        "code",
				Type:        disgord.OptionTypeString,
				Description: "Verify Code",
			},
		},
	}
	verifyStruct = &models.Interaction{
		CommandDefinition: verifyDefinition,
		Callback:          verify,
	}
	verifyFailed = &disgord.CreateInteractionResponse{
		Type: disgord.InteractionCallbackChannelMessageWithSource,
		Data: &disgord.CreateInteractionResponseData{
			Embeds: []*disgord.Embed{
				{
					Title:       "Aww, something went wrong!~",
					Description: "Unknown Error.",
				},
			},
			Flags: disgord.MessageFlagEphemeral,
		},
	}
	verifyDone = &disgord.CreateInteractionResponse{
		Type: disgord.InteractionCallbackChannelMessageWithSource,
		Data: &disgord.CreateInteractionResponseData{
			Embeds: []*disgord.Embed{
				{
					Title:       "Done!",
					Description: "Role given!",
				},
			},
			Flags: disgord.MessageFlagEphemeral,
		},
	}
)

func verify(_ disgord.Session, interactionCreate *disgord.InteractionCreate) {

	answer := verifyFailed
	defer func() {
		discordModels.GetClient().SendInteractionResponse(context.Background(), interactionCreate, answer)
	}()

	if len(interactionCreate.Data.Options) == 0 {
		answer.Data.Embeds[0].Description = "Bad Arguments!"
		return
	}

	db := database.Get()

	code := interactionCreate.Data.Options[0].Value.(string)
	codeData, err := db.GetCodeInfo(code)
	UserID := interactionCreate.Member.UserID.String()

	if err == sql.ErrNoRows {
		answer.Data.Embeds[0].Description = "Code not found :("
		return
	}

	if codeData.Used && codeData.UsedBy != UserID {
		answer.Data.Embeds[0].Description = "Code already used :("
		return
	}

	client := discordModels.GetClient()

	assignRole, err := db.GetRoleID(codeData.AssignRole)

	if err != nil {

		if err == sql.ErrNoRows {
			answer.Data.Embeds[0].Description = "Role not found in database. Please contact administrator!"
		}

		return
	}

	assignRoleSnowflake, _ := disgord.GetSnowflake(assignRole)

	err = client.Guild(interactionCreate.GuildID).Member(interactionCreate.Member.UserID).AddRole(assignRoleSnowflake)
	if err != nil {
		utils.Logger().Println("failed to assign role ->", err)
		answer.Data.Embeds[0].Description = "Failed to assign role :("
		return
	}

	answer = verifyDone
	db.SetUsed(UserID, codeData)
}
