package databaseTables

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	ID        int    `json:"id" gorm:"primaryKey"`
	Username  string `json:"username"`
	DiscordID string `json:"discord_id,omitempty"`
}
