package main

import (
	"fmt"
	"os"
	"strings"
)

var GroupsToCheckFile = "GroupsToCheck.txt"
var GroupsToCheck []string
var GroupsChecked []string

func OpenGroupsToCheckIfExists() {
	_, m := os.Stat(GroupsToCheckFile)
	if m != nil {
		f, e := os.Create(GroupsToCheckFile)
		CheckForErr(e)
		CheckForErr(f.Close())
	}

	f, e := os.ReadFile(GroupsToCheckFile)
	CheckForErr(e)

	GroupsToCheck = strings.Split(string(f), "\n")
	for i, g := range GroupsToCheck {
		GroupsToCheck[i] = strings.Trim(g, " ")
	}
}

func SaveGroupsToCheckToFile() {
	f, e := os.OpenFile(GroupsToCheckFile, os.O_WRONLY, os.ModePerm)
	CheckForErr(e)
	for i, g := range GroupsToCheck {
		_, e = f.WriteString(g)
		if i < len(GroupsToCheck)-1 {
			f.WriteString("\n")
		}
		CheckForErr(e)

	}
}

var UserIDs []string

func StartGroupSearch() {
	OpenGroupsToCheckIfExists()

	groupCount := len(GroupsToCheck)

	for i := 0; i < groupCount; i++ {

		CheckGroup(i)
	}
}

func CheckGroup(index int) []GroupMember {
	group := GroupsToCheck[index]
	fmt.Println(group)
	JoinGroup(group)
	//SaveGroupsToCheckToFile()
	members := GetGroupMembers(group)
	return members
}
