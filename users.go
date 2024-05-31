package main

import (
	"fmt"
	"time"
)

type Badge struct {
	AssignedAt       time.Time `json:"assignedAt"`
	BadgeDescription string    `json:"badgeDescription"`
	BadgeId          string    `json:"badgeId"`
	BadgeImageUrl    string    `json:"badgeImageUrl"`
	BadgeName        string    `json:"badgeName"`
	Hidden           bool      `json:"hidden"`
	Showcased        bool      `json:"showcased"`
	UpdatedAt        time.Time `json:"updatedAt"`
}

type User struct {
	AllowAvatarCopying             bool     `json:"allowAvatarCopying"`
	Badges                         []Badge  `json:"badges"`
	Bio                            string   `json:"bio"`
	BioLinks                       []string `json:"bioLinks"`
	CurrentAvatarImageUrl          string   `json:"currentAvatarImageUrl"`
	CurrentAvatarThumbnailImageUrl string   `json:"currentAvatarThumbnailImageUrl"`
	CurrentAvatarTags              []string `json:"currentAvatarTags"`
	DateJoined                     string   `json:"date_joined"`
	DeveloperType                  string   `json:"developerType"`
	DisplayName                    string   `json:"displayName"`
	FriendKey                      string   `json:"friendKey"`
	FriendRequestStatus            string   `json:"friendRequestStatus"`
	Id                             string   `json:"id"`
	InstanceId                     string   `json:"instanceId"`
	IsFriend                       bool     `json:"isFriend"`
	LastActivity                   string   `json:"last_activity"`
	LastLogin                      string   `json:"last_login"`
	LastPlatform                   string   `json:"last_platform"`
	Location                       string   `json:"location"`
	Note                           string   `json:"note"`
	ProfilePicOverride             string   `json:"profilePicOverride"`
	Pronouns                       string   `json:"pronouns"`
	State                          string   `json:"state"`
	Status                         string   `json:"status"`
	StatusDescription              string   `json:"statusDescription"`
	Tags                           []string `json:"tags"`
	TravelingToInstance            string   `json:"travelingToInstance"`
	TravelingToLocation            string   `json:"travelingToLocation"`
	TravelingToWorld               string   `json:"travelingToWorld"`
	UserIcon                       string   `json:"userIcon"`
	WorldId                        string   `json:"worldId"`
}

var LocalUsersToCheck IDList

func StartUserParser() {

	for {
		time.Sleep(time.Second)
		UserIDs.mu.Lock()
		fmt.Println(len(UserIDs.ids))
		LocalUsersToCheck.mu.Lock()
		LocalUsersToCheck.ids = append(LocalUsersToCheck.ids, UserIDs.ids...)
		LocalUsersToCheck.mu.Unlock()

		UserIDs.ids = []string{}

		UserIDs.mu.Unlock()
		if len(UserIDs.ids) > 0 {
			CheckUsers()
		}
	}
}

func CheckUsers() {

}
