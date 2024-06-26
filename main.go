package main

import (
	"fmt"

	"github.com/gminato/first-discord-bot/bot"
	"github.com/gminato/first-discord-bot/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()
	<-make(chan struct{})
	return
}
