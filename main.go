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
	if m == nil {
		return
	}
	c, e := os.Create(BotFileName)
	CheckForErr(e)
	defer c.Close()

	bUser := BotUser{}
	bUsers := []BotUser{bUser}
	bUserJson, _ := json.Marshal(bUsers)

	_, e = c.Write(bUserJson)
	CheckForErr(e)
	fmt.Println("Bots.json was not found, so one was created, please fill in bot info and relaunch application")
	os.Exit(1)

}

func SaveBotJson() {
	b := make([]byte, 1000)
	b, err := json.Marshal(BotUsers)
	CheckForErr(err)
	f, e := os.OpenFile(BotFileName, os.O_WRONLY, os.ModePerm)
	CheckForErr(e)
	_, e = f.Write(b)
	CheckForErr(e)
	f.Close()
}

func main() {
	CheckIfBotJsonExists()
	bUser, err := os.ReadFile(BotFileName)
	CheckForErr(err)
	err = json.Unmarshal(bUser, &BotUsers)
	CheckForErr(err)
	CreateClient()
	LogInBots()

	StartGroupSearch()
}

func LogInBots() {
	for i := 0; i < len(BotUsers); i++ {
		cBot := BotUsers[i]
		if len(cBot.AuthCookie) > 1 {
			SetBotCookie(i)
			valid := CheckIfCookieIsValid()
			fmt.Printf("Bot %d's cookie check result: %t\n", i, valid)
			if valid {
				continue
			}
		}
		fmt.Printf("Bot %d's has no cookie, or it is invalid\n", i)

		c, otp := Login(cBot.Username, cBot.Password)
		BotUsers[i].AuthCookie = c
		SaveBotJson()
		if otp {
			//implement joe api
			fmt.Println("Need otp")
		}
	}
}
