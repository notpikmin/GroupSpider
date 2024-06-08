package main

import (
	"github.com/gtuk/discordwebhook"
	"strconv"
	"time"
)

var username = "BioCreeper"

func CreateEmbed(user User, score int) *discordwebhook.Embed {
	url := "https://vrchat.com/home/user/" + user.Id
	description := "No status"
	if user.StatusDescription != "" {
		description = user.StatusDescription
	}
	color := "512"
	image := user.CurrentAvatarImageUrl
	if user.ProfilePicOverride != "" {
		image = user.ProfilePicOverride
	}
	userIcon := image
	if user.UserIcon != "" {
		userIcon = user.UserIcon
	}
	author := discordwebhook.Author{
		Name:    &user.DisplayName,
		Url:     &url,
		IconUrl: &userIcon,
	}
	bio := "Bio"
	scoreName := "Score"
	scoreString := strconv.Itoa(score)

	fields := []discordwebhook.Field{
		{
			Name:  &scoreName,
			Value: &scoreString,
		}, {
			Name:  &bio,
			Value: &user.Bio,
		}}

	embed := discordwebhook.Embed{
		Title:       &user.DisplayName,
		Url:         &url,
		Description: &description,
		Color:       &color,
		Author:      &author,
		Fields:      &fields,
		Thumbnail:   nil,
		Image:       &discordwebhook.Image{Url: &image},
		Footer:      nil,
	}
	return &embed
}

func SendEmbed(user User, score int) {
	var url = "https://discord.com/api/webhooks/1005687599286976522/6Ferdnfxy4sBFYhMWxqTQDWjJDxaueq2QuyiCQVA0G1NnnZKoCf_-Mk6eLcW8qCha8Bo"
	//var url = "https://discord.com/api/webhooks/1246980704370163724/lJAKIEzicMkPhynEx7SAoEFxKc5MDNcSUTkGI-rFPX1B041_oqLvdUTAF6888Yu9x3FI"

	embed := CreateEmbed(user, score)
	message := discordwebhook.Message{
		Username: &username,
		Embeds:   &[]discordwebhook.Embed{*embed},
	}

	err := discordwebhook.SendMessage(url, message)

	CheckForErr(err)
	time.Sleep(2 * time.Second)
	Send(user, score)

}
func Send(user User, score int) {
	var url = "https://discord.com/api/webhooks/1248333994421125223/hxmZPzrTIgu6mGIZ1FXcaeGjx61qGJ5ksJL53HnuaOM0k5038vllLIwh6_NOw8Id0dLD"

	embed := CreateEmbed(user, score)
	message := discordwebhook.Message{
		Username: &username,
		Embeds:   &[]discordwebhook.Embed{*embed},
	}

	err := discordwebhook.SendMessage(url, message)

	CheckForErr(err)

}
