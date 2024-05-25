package main

import (
	"encoding/json"
	"fmt"
	"strconv"
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
	var b []byte
	c := 5
	for i := 0; i < 1; i++ {
		res := MakeRequest(BaseApi+"groups/"+id+"/members?n="+strconv.Itoa(c)+"&offset="+strconv.Itoa(o), "GET", "", nil)

		o += c

		fmt.Println(string(b))
		tMember := make([]GroupMember, c)
		err := json.NewDecoder(res.Body).Decode(&tMember)
		CheckForErr(err)
		members = append(members, tMember...)
		//time.Sleep(500 * time.Millisecond)
	}
	return members
}
