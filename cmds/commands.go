package cmds

import (
	"github.com/bwmarrin/discordgo"
)

type Command struct {
	Command *discordgo.ApplicationCommand
	Action  func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

var Cmds = []Command{
	{
		Command: &discordgo.ApplicationCommand{
			Name:        "ping",
			Description: "Ping-pong",
		},
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
