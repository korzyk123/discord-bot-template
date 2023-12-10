package cmds

import (
	"github.com/bwmarrin/discordgo"
)

type CommandDefinition struct {
	Command *discordgo.ApplicationCommand
	Action  func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

// Cmds is a variable that holds a slice of CommandDefinition structs. Each CommandDefinition struct represents an application command for a Discord bot.
// CommandDefinition struct contains two fields:
// -- Command, which is a pointer to a discordgo.ApplicationCommand struct representing the details of the command, such as its name and description.
// -- Action, which is a function that takes a discordgo.Session and discordgo.InteractionCreate as parameters and defines the action to be performed when the command is invoked.
// Example usage of the Cmds variable can be found in the README of this repo.
var Cmds = []CommandDefinition{
	{
		Command: &discordgo.ApplicationCommand{
			Name:        "",
			Description: "",
		},
		Action: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			// Your code goes here
		},
	},
}
