package main

import (
	"flag"
	"github.com/bwmarrin/discordgo"
	"log"
	"manifest/cmds"
	"os"
	"os/signal"
)

var (
	AppID string
	Token string
)

func main() {

	flag.StringVar(&AppID, "i", "", "Application ID")
	flag.StringVar(&Token, "t", "", "Bot token")
	flag.Parse()

	s, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Fatal(err)
		return
	}

	ref := cmds.Cmds

	commands := make([]*discordgo.ApplicationCommand, len(ref))

	for i, c := range ref {
		commands[i] = &discordgo.ApplicationCommand{
			Name:        c.Name,
			Description: c.Description,
		}
	}

	_, err = s.ApplicationCommandBulkOverwrite(AppID, "", commands)
	if err != nil {
		log.Fatal(err)
		return
	}

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type != discordgo.InteractionApplicationCommand {
			return
		}

		data := i.ApplicationCommandData()

		for _, c := range ref {
			if c.Name == data.Name {
				c.Action(s, i)
			}
		}
	})

	err = s.Open()
	if err != nil {
		log.Fatal(err)
		return
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	err = s.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
}
