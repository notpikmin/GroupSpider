package main

import "encoding/json"

func GetUser(id string) User {
	res := MakeRequest(BaseApi+"users/"+id, "GET", "", nil)
	var user User
	err := json.NewDecoder(res.Body).Decode(&user)
	CheckForErr(err)

	return user
}
