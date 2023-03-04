package interactions

import (
	"context"
	"database/sql"
	"github.com/yoonaowo/discord_verifier/internal/translations"

	"github.com/andersfylling/disgord"

	"github.com/yoonaowo/discord_verifier/internal/database"
	"github.com/yoonaowo/discord_verifier/internal/models"
	discordModels "github.com/yoonaowo/discord_verifier/internal/models/discord"
	"github.com/yoonaowo/discord_verifier/internal/utils"
)

var (
	verifyDefinition = &disgord.CreateApplicationCommand{
		Name:        "verify",
		Description: translations.Get("VERIFY_DEF_DESCRIPTION"),
		Options: []*disgord.ApplicationCommandOption{
			{
				Required:    true,
				Name:        translations.Get("VERIFY_DEF_OPT_NAME"),
				Type:        disgord.OptionTypeString,
				Description: translations.Get("VERIFY_DEF_OPT_DESCRIPTION"),
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
					Title:       translations.Get("SOMETHING_WENT_WRONG"),
					Description: translations.Get("UNK_ERROR"),
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
					Title:       translations.Get("VERIFY_DONE_TITLE"),
					Description: translations.Get("VERIFY_DONE_DESCRIPTION"),
				},
			},
			Flags: disgord.MessageFlagEphemeral,
		},
	}
)

func verify(_ disgord.Session, interactionCreate *disgord.InteractionCreate) {

	answer := verifyFailed
	defer func() {
		_ = discordModels.GetClient().SendInteractionResponse(context.Background(), interactionCreate, answer)
	}()

	if len(interactionCreate.Data.Options) == 0 {
		answer.Data.Embeds[0].Description = translations.Get("BAD_ARGUMENTS") // never reached?
		return
	}

	db := database.Get()

	code := interactionCreate.Data.Options[0].Value.(string)
	codeData, err := db.GetCodeInfo(code)
	UserID := interactionCreate.Member.UserID.String()

	if err == sql.ErrNoRows {
		answer.Data.Embeds[0].Description = translations.Get("CODE_NOT_FOUND")
		return
	}

	if codeData.Used && codeData.UsedBy != UserID {
		answer.Data.Embeds[0].Description = translations.Get("CODE_ALREADY_USED")
		return
	}

	client := discordModels.GetClient()

	assignRole, err := db.GetRoleID(codeData.AssignRole)

	if err != nil {

		if err == sql.ErrNoRows {
			answer.Data.Embeds[0].Description = translations.Get("ROLE_NOT_FOUND_CONTACT_ADMIN")
		}

		return
	}

	assignRoleSnowflake, _ := disgord.GetSnowflake(assignRole)

	err = client.Guild(interactionCreate.GuildID).Member(interactionCreate.Member.UserID).AddRole(assignRoleSnowflake)
	if err != nil {
		utils.Logger().Println("failed to assign role ->", err)
		answer.Data.Embeds[0].Description = translations.Get("FAILED_TO_ASSIGN_ROLE")
		return
	}

	answer = verifyDone
	db.SetUsed(UserID, codeData)
}
