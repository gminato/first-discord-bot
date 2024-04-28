package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/gminato/first-discord-bot/config"
)

var BotId string
var goBot *discordgo.Session

func Start() {

	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	BotId = u.ID

	goBot.AddHandler(messageHandle)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running")
}

func messageHandle(s *discordgo.Session, m *discordgo.MessageCreate) {
	if BotId == m.Author.ID {
		return
	}
	if m.Content == "ping" {
		_, err := s.ChannelMessageSend(m.ChannelID, "pong")

		if err != nil {
			fmt.Print(err.Error())
		}
	}
}
