package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"sync"
	"time"
)

type IDList struct {
	mu  sync.Mutex
	ids []string
}

var GroupsToCheckFile = "GroupsToCheck.txt"
var GroupsToCheck IDList
var GroupsToAddToCheckList IDList
var GroupsChecked IDList
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

func AddGroupToCheckId(id string) {
	if slices.Contains(GroupsChecked.ids, id) {
		return
	}
	GroupsChecked.mu.Lock()
	GroupsToAddToCheckList.mu.Lock()
	GroupsChecked.ids = append(GroupsChecked.ids, id)
	GroupsToAddToCheckList.ids = append(GroupsToAddToCheckList.ids, id)

	GroupsToAddToCheckList.mu.Unlock()
	GroupsChecked.mu.Unlock()

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
	for {
		GroupsToCheck.mu.Lock()
		for i := 0; i < len(GroupsToCheck.ids); i++ {
			UserIDs.mu.Lock()
			members := CheckGroup(i)
			fmt.Println(len(members))
			for x := 0; x < len(members); x++ {
				UserIDs.ids = append(UserIDs.ids, members[x].UserID)
			}
			UserIDs.mu.Unlock()
			fmt.Printf("Finished Checking group: %s\n", GroupsToCheck.ids[i])
		}
		GroupsToAddToCheckList.mu.Lock()
		GroupsToCheck.ids = GroupsToAddToCheckList.ids
		GroupsToAddToCheckList.ids = []string{}

		GroupsToAddToCheckList.mu.Unlock()
		GroupsToCheck.mu.Unlock()
		time.Sleep(2 * time.Second)
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
