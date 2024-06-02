package main

import (
	"github.com/gtuk/discordwebhook"
)

var username = "BioCreeper"

func CreateEmbed(user User) *discordwebhook.Embed {
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
	bio := "bio"
	fields := []discordwebhook.Field{{
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

func SendEmbed(user User) {
	var url = "https://discord.com/api/webhooks/1005687599286976522/6Ferdnfxy4sBFYhMWxqTQDWjJDxaueq2QuyiCQVA0G1NnnZKoCf_-Mk6eLcW8qCha8Bo"
	embed := CreateEmbed(user)
	message := discordwebhook.Message{
		Username: &username,
		Embeds:   &[]discordwebhook.Embed{*embed},
	}

	err := discordwebhook.SendMessage(url, message)
	CheckForErr(err)
}
