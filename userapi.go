package main

import (
	"encoding/json"
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
type UserGroup struct {
	Id                string    `json:"id"`
	Name              string    `json:"name"`
	ShortCode         string    `json:"shortCode"`
	Discriminator     string    `json:"discriminator"`
	Description       string    `json:"description"`
	IconId            string    `json:"iconId"`
	IconUrl           string    `json:"iconUrl"`
	BannerId          string    `json:"bannerId"`
	BannerUrl         string    `json:"bannerUrl"`
	Privacy           string    `json:"privacy"`
	LastPostCreatedAt time.Time `json:"lastPostCreatedAt"`
	OwnerId           string    `json:"ownerId"`
	MemberCount       int       `json:"memberCount"`
	GroupId           string    `json:"groupId"`
	MemberVisibility  string    `json:"memberVisibility"`
	IsRepresenting    bool      `json:"isRepresenting"`
	MutualGroup       bool      `json:"mutualGroup"`
	LastPostReadAt    time.Time `json:"lastPostReadAt"`
}

func GetUser(id string) User {
	res := MakeRequest(BaseApi+"users/"+id, "GET", "", nil)
	var user User
	err := json.NewDecoder(res.Body).Decode(&user)
	CheckForErr(err)

	return user
}

func GetUsersGroups(id string) []string {
	res := MakeRequest(BaseApi+"users/"+id+"/groups", "GET", "", nil)
	var ug []UserGroup
	err := json.NewDecoder(res.Body).Decode(&ug)
	CheckForErr(err)
	groups := []string{}
	for _, g := range ug {
		groups = append(groups, g.GroupId)
	}
	return groups
}
