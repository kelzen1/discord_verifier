package utils

import "os"

var (
	FlagAdminPassword  = os.Getenv("ADMIN_PASSWORD")
	FlagDiscordGuild   = os.Getenv("DISCORD_GUILD")
	FlagDiscordToken   = os.Getenv("DISCORD_TOKEN")
	FlagSecretPassword = os.Getenv("SECRET")
	FlagWebPort 	   = os.Getenv("PORT")
)
