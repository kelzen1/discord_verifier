package interactions

import (
	"Verifier/database"
	"Verifier/database/actions"
	"Verifier/models"
	databaseTables "Verifier/models/database"
	discordModels "Verifier/models/discord"
	"context"
	"database/sql"
	"github.com/andersfylling/disgord"
	"log"
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
		verifyFailed.Data.Embeds[0].Description = "Bad Arguments!"
		return
	}

	code := interactionCreate.Data.Options[0].Value.(string)
	codeData, err := actions.GetCodeInfo(code)
	UserID := interactionCreate.Member.UserID.String()
	if err != nil || (codeData.Used && codeData.UsedBy != UserID) {
		verifyFailed.Data.Embeds[0].Description = "Code already used!"
		return
	}

	client := discordModels.GetClient()

	assignRole, err := actions.GetRoleID(codeData.AssignRole)

	if err != nil {
		return
	}

	assignRoleSnowflake, _ := disgord.GetSnowflake(assignRole)

	err = client.Guild(interactionCreate.GuildID).Member(interactionCreate.Member.UserID).AddRole(assignRoleSnowflake)
	if err != nil {
		log.Println("failed to assign role ->", err)
		verifyFailed.Data.Embeds[0].Description = "Failed to assign role :("
		return
	}

	answer = verifyDone

	db := database.Get()

	updateData := map[string]interface{}{
		"used":    true,
		"used_by": UserID,
	}

	db.Table("codes").Where("code = ?", codeData.Code).Updates(updateData)

	if codeData.Username == "unknown" {
		return
	}

	db = db.Table("users")

	findQuery := db.Where("username = ? AND discord_id = ?", codeData.Username, UserID)

	_, err = database.ScanMany[databaseTables.Users](findQuery)

	if err == sql.ErrNoRows {
		db.Create(&databaseTables.Users{
			Username:  codeData.Username,
			DiscordID: UserID,
		})
	}
}
