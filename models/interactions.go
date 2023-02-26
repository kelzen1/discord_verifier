package models

import "github.com/andersfylling/disgord"

type Interaction struct {
	Callback          func(s disgord.Session, h *disgord.InteractionCreate)
	CommandDefinition *disgord.CreateApplicationCommand
}
