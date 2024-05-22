package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type BotUser struct {
	Username            string
	Password            string
	AuthCookie          string
	TwoFactorAuthCookie string
	Proxy               string
}

var BotUsers []BotUser

var BotFileName = "Bots.json"

func CheckIfBotJsonExists() {
	_, m := os.Stat(BotFileName)
	if m != nil {
		c, e := os.Create(BotFileName)
		CheckForErr(e)
		defer c.Close()

		bUser := BotUser{}
		bUsers := []BotUser{bUser}
		bUserJson, _ := json.Marshal(bUsers)

		c.Write(bUserJson)

		fmt.Println("Bots.json was not found, so one was created, please fill in bot info and relaunch application")
		os.Exit(1)
	}
}

func main() {
	CheckIfBotJsonExists()

	bUser, err := os.ReadFile(BotFileName)
	CheckForErr(err)
	err = json.Unmarshal(bUser, &BotUsers)
	CheckForErr(err)
	CreateClient()
	CheckIfCookieIsValid()
}
