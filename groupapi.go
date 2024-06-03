package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

func JoinGroup(id string) {

	res := MakeRequest(BaseApi+"groups/"+id+"/join", "POST", "", nil)

	switch res.StatusCode {
	case 200:
		fmt.Println("Joined group: " + id + ",successfully")
		break
	case 400:
		fmt.Println("Already A member of: " + id)
		break
	case 401:
		fmt.Println("Not logged in group join failed")
		break
	case 404:
		fmt.Println("Group:" + id + ", not found")
		break
	default:
		fmt.Println("Unknown status code:" + strconv.Itoa(res.StatusCode))

	}
}

func GetGroupMembers(id string) []GroupMember {
	var members []GroupMember
	o := 0
	c := 100
	for i := 0; i < 1000; i++ {
		res := MakeRequest(BaseApi+"groups/"+id+"/members?n="+strconv.Itoa(c)+"&offset="+strconv.Itoa(o), "GET", "", nil)

		o += c

		tMember := make([]GroupMember, c)
		err := json.NewDecoder(res.Body).Decode(&tMember)

		if CheckForErr(err) || len(tMember) < 5 {
			break
		}
		members = append(members, tMember...)
		time.Sleep(2 * time.Second)

	}

	return members
}
