package main

import "fmt"

func CheckIfCookieIsValid() {
	AddCookie("auth", BotUsers[0].AuthCookie)
	AddCookie("twoFactorAuth", BotUsers[0].TwoFactorAuthCookie)

	res := DoRequest(MakeRequest("https://api.vrchat.cloud/api/1/auth", "GET"))
	b := make([]byte, 1000)
	res.Body.Read(b)
	fmt.Println(string(b))
}
