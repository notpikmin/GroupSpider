package main

import "time"

type GroupUser struct {
	ID                             string   `json:"id"`
	DisplayName                    string   `json:"displayName"`
	ThumbnailURL                   string   `json:"thumbnailUrl"`
	IconURL                        string   `json:"iconUrl"`
	ProfilePicOverride             string   `json:"profilePicOverride"`
	CurrentAvatarThumbnailImageURL string   `json:"currentAvatarThumbnailImageUrl"`
	CurrentAvatarTags              []string `json:"currentAvatarTags"`
}
type GroupMember struct {
	ID                          string    `json:"id"`
	GroupID                     string    `json:"groupId"`
	UserID                      string    `json:"userId"`
	IsRepresenting              bool      `json:"isRepresenting"`
	User                        GroupUser `json:"user"`
	RoleIds                     []string  `json:"roleIds"`
	MRoleIds                    []string  `json:"mRoleIds"`
	JoinedAt                    time.Time `json:"joinedAt"`
	MembershipStatus            string    `json:"membershipStatus"`
	Visibility                  string    `json:"visibility"`
	IsSubscribedToAnnouncements bool      `json:"isSubscribedToAnnouncements"`
	CreatedAt                   time.Time `json:"createdAt"`
	BannedAt                    time.Time `json:"bannedAt"`
	ManagerNotes                string    `json:"managerNotes"`
	LastPostReadAt              time.Time `json:"lastPostReadAt"`
	HasJoinedFromPurchase       bool      `json:"hasJoinedFromPurchase"`
}
