package databaseTables

type Users struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	DiscordID string `json:"discord_id,omitempty"`
}
