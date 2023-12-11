# dbt `discord-bot-template`

> discord bot template with slash command support

![GitHub License](https://img.shields.io/github/license/korzyk123/discord-bot-template?style=flat-square)

Discord bot template with support for custom slash commands written in Go

## Getting Started

### Prerequisites

You need to have Go installed on your local machine. This project uses Go 1.21.

### Installation

To start, clone this repo

```bash
git clone https://github.com/korzyk123/discord-bot-template
```

navigate to cloned dir

```bash
cd discord-bot-template
```

### Running

Run the main package

`-i` application ID<br>
`-t` bot token

```bash
go run discord-bot-template -i appid -t token
```

## Usage

All slash commands are defined in `cmds/commands.go` using the built-in `CommandDefinition` struct

Examples:

A command with interaction response
```go
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
}
```

A command with one required argument
```go
{
    Command: &discordgo.ApplicationCommand{
        Name:        "reply",
        Description: "Reply with custom message",
        Options: []*discordgo.ApplicationCommandOption{
            {
                Name:        "msg",
                Description: "Message",
                Type:        discordgo.ApplicationCommandOptionString,
                Required:    true,
            },
        },
    },
    Action: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
        _ = s.InteractionRespond(
            i.Interaction,
            &discordgo.InteractionResponse{
                Type: discordgo.InteractionResponseChannelMessageWithSource,
                Data: &discordgo.InteractionResponseData{
                    Content: i.ApplicationCommandData().Options[0].StringValue(),
                },
            },
        )
    },
}
```

A command with multiple arguments, one of which is optional
```go
{
    Command: &discordgo.ApplicationCommand{
        Name:        "pow",
        Description: "Return result of exponentiation",
        Options: []*discordgo.ApplicationCommandOption{
            {
                Name:        "base",
                Description: "Base",
                Type:        discordgo.ApplicationCommandOptionInteger,
                Required:    true,
            },
            {
                Name:        "exp",
                Description: "Exponent",
                Type:        discordgo.ApplicationCommandOptionInteger,
                Required:    false,
            },
        },
    },
    Action: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
        var args []float64

        for _, arg := range i.ApplicationCommandData().Options {
            args = append(args, float64(arg.IntValue()))
        }

        if len(args) < 2 {
            args = append(args, 2) // Default value for 2nd argument
        }

        result := int64(math.Pow(args[0], args[1]))

        _ = s.InteractionRespond(
            i.Interaction,
            &discordgo.InteractionResponse{
                Type: discordgo.InteractionResponseChannelMessageWithSource,
                Data: &discordgo.InteractionResponseData{
                    Content: strconv.FormatInt(result, 10),
                },
            },
        )
    },
}
```
Whenever the package starts, the slash commands are automatically overwritten in bulk to match the local code.

If you modify the source code, you might need to stop the running instance and `go run` it again to apply the changes to production.

## Built With

- [discordgo](https://github.com/bwmarrin/discordgo) - the discord library used

## License

This project is licensed under the Unlicense license - see the [LICENSE](LICENSE) file for details