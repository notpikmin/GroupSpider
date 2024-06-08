package main

import (
	"fmt"
	"slices"
	"time"
)

var LocalUsersToCheck IDList
var UsersChecked IDList

func StartUserParser() {

	for {
		time.Sleep(3 * time.Second)
		UserIDs.mu.Lock()
		LocalUsersToCheck.mu.Lock()
		LocalUsersToCheck.ids = append(LocalUsersToCheck.ids, UserIDs.ids...)
		LocalUsersToCheck.mu.Unlock()
		fmt.Printf("local: %d \n", len(LocalUsersToCheck.ids))

		fmt.Printf("user: %d \n", len(UserIDs.ids))

		UserIDs.ids = []string{}

		UserIDs.mu.Unlock()
		if len(LocalUsersToCheck.ids) > 0 {
			CheckUsers()
		}
	}
}

func CheckUsers() {
	LocalUsersToCheck.mu.Lock()

	for i := 0; i < len(LocalUsersToCheck.ids); i++ {
		UsersChecked.mu.Lock()
		userID := LocalUsersToCheck.ids[i]
		if slices.Contains(UsersChecked.ids, userID) {
			fmt.Println("User already checked")
			continue
		}
		UsersChecked.ids = append(UsersChecked.ids, userID)
		u := GetUser(userID)
		HandleUser(u)

		UsersChecked.mu.Unlock()

		time.Sleep(3 * time.Second)
	}
	LocalUsersToCheck.ids = []string{}
	LocalUsersToCheck.mu.Unlock()

}

func HandleUser(user User) {
	groups := GetUsersGroups(user.Id)
	for _, group := range groups {
		AddGroupToCheckId(group)
	}
	score := CringeRate(&user)

	if score != 0 {
		SendEmbed(user, score)
	}

}
