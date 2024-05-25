package main

type GroupMember struct {
	Id             string    `json:"id"`
	GroupId        string    `json:"groupId"`
	UserId         string    `json:"userId"`
	IsRepresenting bool      `json:"isRepresenting"`
	User           GroupUser `json:"user"`
}

// not full object
type GroupUser struct {
	Id          string `json:"id"`
	DisplayName string `json:"displayName"`
	Visibility  string `json:"visibility"`
}
