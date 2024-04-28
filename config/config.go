package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

var (
	Token     string
	BotPrefix string

	config *configStruct
)

type configStruct struct {
	Token     string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
}

func ReadConfig() error {
	fmt.Println("Reading config files")
	// file, err := io.ReadAll(os.Open("./config.json"))
	filePointer, err := os.Open("./config.json")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	// fmt.Println(*filePointer)

	file, err := io.ReadAll(filePointer)

	if err != nil {
		fmt.Println((err.Error()))
		return err
	}

	if err != nil {
		fmt.Println((err.Error()))
		return err
	}

	fmt.Println(string(file))

	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	Token = config.Token
	BotPrefix = config.BotPrefix
	return nil
}
