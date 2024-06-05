package main

import (
	"bytes"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

var Client http.Client
var VRCUrl *url.URL

func CreateClient() {
	jar, err := cookiejar.New(nil)
	var cookies []*http.Cookie
	VRCUrl, _ = url.Parse("https://vrchat.com")

	jar.SetCookies(VRCUrl, cookies)
	CheckForErr(err)
	Client = http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           jar,
		Timeout:       0,
	}

}
func AddCookie(name string, value string) {
	authCookie := &http.Cookie{
		Name:   name,
		Value:  value,
		Domain: "vrchat.com",
		Path:   "/",
	}

	Client.Jar.SetCookies(VRCUrl, append(Client.Jar.Cookies(VRCUrl), authCookie))
}

func MakeRequest(reqUrl string, method string, body string, headers map[string]string) *http.Response {
	jsonBody := []byte(body)
	bodyReader := bytes.NewReader(jsonBody)
	req, err := http.NewRequest(method, reqUrl, bodyReader)
	CheckForErr(err)
	req.Header = map[string][]string{
		"Accept":     {"*/*"},
		"Host":       {"vrchat.com"},
		"User-Agent": {"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:126.0) Gecko/20100101 Firefox/126.0"},
	}

	for k, v := range headers {
		req.Header[k] = []string{v}
	}
	return doRequest(req)
}

func doRequest(req *http.Request) *http.Response {

	cs := Client.Jar.Cookies(VRCUrl)
	for _, c := range cs {
		req.AddCookie(c)
	}
	response, err := Client.Do(req)

	CheckForErr(err)

	return response
}
