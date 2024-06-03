package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var CringeMetric []CringeItem

type CringeItem struct {
	Key   string
	Score int
}

func CringeRate(user User) int {
	cringeScore := 0

	bio := user.Bio

	for _, keyword := range CringeMetric {
		c := strings.Count(bio, keyword.Key)
		cringeScore += c * keyword.Score
	}

	return cringeScore
}

func DownloadCringeMetric() {

	fullURLFile := "https://rentry.co/cringemetric/raw"
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	// Put content on file
	resp, err := client.Get(fullURLFile)
	CheckForErr(err)
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	response := strings.Split(string(b), "\n")
	for _, s := range response {
		ci := strings.Split(s, ":")
		if len(ci) < 2 {
			continue
		}
		score, err := strconv.Atoi(strings.Trim(ci[1], "\r"))

		CheckForErr(err)
		cringe := CringeItem{
			Key:   ci[0],
			Score: score,
		}
		CringeMetric = append(CringeMetric, cringe)
	}
}
