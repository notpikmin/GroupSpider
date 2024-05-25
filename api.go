package main

import (
	"encoding/base64"
	"fmt"
	"strings"
)

var BaseApi = "https://api.vrchat.cloud/api/1/"

func SetBotCookie(botIndex int) {
	AddCookie("auth", BotUsers[botIndex].AuthCookie)
	AddCookie("twoFactorAuth", BotUsers[botIndex].TwoFactorAuthCookie)
}

func CheckIfCookieIsValid() bool {
	res := MakeRequest(BaseApi+"auth", "GET", "", nil)
	return res.StatusCode == 200
}

func BasicAuth(username, password string) string {
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(username+":"+password))
}

func Login(username, password string) (string, bool) {
	res := MakeRequest(BaseApi+"auth/user", "GET", "", map[string]string{"Authorization": BasicAuth(username, password)})
	cookies := res.Cookies()

	b := make([]byte, 1000)
	res.Body.Read(b)
	fmt.Println(string(b))
	needOtp := strings.Contains(string(b), "Otp")
	fmt.Println(cookies[0].Value)
	return cookies[0].Value, needOtp
}
