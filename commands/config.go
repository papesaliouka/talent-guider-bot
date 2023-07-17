package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	Token       string
	BotPrefix   string
	GuiderToken string
	guildID     string
	config      *configStruct
)

type configStruct struct {
	Token       string `json:"Token"`
	BotPrefix   string `json:"BotPrefix"`
	GuiderToken string `json:"GuiderToken"`
	guildID     string `json:"GuildID"`
}

func ReadConfig() error {
	fmt.Println("Reading config file...")
	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		fmt.Println(err.Error())
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
	GuiderToken = config.GuiderToken
	guildID = config.guildID

	return nil
}
