package cmds

import (
	"github.com/bwmarrin/discordgo"
)

type Command struct {
	Name        string
	Description string
	Action      func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

var Cmds = []Command{
	{
		Name:        "ping",
		Description: "Ping-pong",
		Action: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			_ = s.InteractionRespond(
				i.Interaction,
				&discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "pong",
					},
				},
			)
		},
	},
}
