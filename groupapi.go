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
	b := make([]byte, 10000)

	for i := 0; i < 1; i++ {
		res := MakeRequest(BaseApi+"groups/"+id+"/members?n=5&offset="+strconv.Itoa(o), "GET", "", nil)

		o += 100

		res.Body.Read(b)
		if len(b) < 8 {
			break
		}
		fmt.Println(string(b))
		tMember := make([]GroupMember, 100)
		//fmt.Println(string(b))
		err := json.Unmarshal(b, &tMember)
		CheckForErr(err)
		members = append(members, tMember...)
		time.Sleep(500 * time.Millisecond)
	}
	return members
}
