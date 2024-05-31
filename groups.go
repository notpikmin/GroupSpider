package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
)

type IDList struct {
	mu  sync.Mutex
	ids []string
}

var GroupsToCheckFile = "GroupsToCheck.txt"
var GroupsToCheck IDList

var UserIDs IDList

func OpenGroupsToCheckIfExists() {
	_, m := os.Stat(GroupsToCheckFile)
	if m != nil {
		f, e := os.Create(GroupsToCheckFile)
		CheckForErr(e)
		CheckForErr(f.Close())
	}

	f, e := os.ReadFile(GroupsToCheckFile)
	CheckForErr(e)
	GroupsToCheck.mu.Lock()
	GroupsToCheck.ids = strings.Split(string(f), "\n")
	for i, g := range GroupsToCheck.ids {
		GroupsToCheck.ids[i] = strings.Trim(g, " ")
	}
	GroupsToCheck.mu.Unlock()

}

func SaveGroupsToCheckToFile() {
	f, e := os.OpenFile(GroupsToCheckFile, os.O_WRONLY, os.ModePerm)
	CheckForErr(e)
	for i, g := range GroupsToCheck.ids {
		_, e = f.WriteString(g)
		if i < len(GroupsToCheck.ids)-1 {
			f.WriteString("\n")
		}
		CheckForErr(e)

	}
}

func StartGroupSearch() {
	OpenGroupsToCheckIfExists()
	groupCount := len(GroupsToCheck.ids)

	for i := 0; i < groupCount; i++ {
		UserIDs.mu.Lock()
		members := CheckGroup(i)

		for x := 0; x < len(members); x++ {
			UserIDs.ids = append(UserIDs.ids, members[x].UserID)
		}
		UserIDs.mu.Unlock()

	}
}

func CheckGroup(index int) []GroupMember {
	group := GroupsToCheck.ids[index]
	fmt.Println(group)
	JoinGroup(group)
	//SaveGroupsToCheckToFile()
	members := GetGroupMembers(group)
	return members
}
