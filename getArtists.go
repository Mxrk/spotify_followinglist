package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func getArtists() {
	// URL for the api endpoint
	URL := "https://api.spotify.com/v1/me/following?"
	reqType := "&type=artist"

	// default limit is 20 (parameter=limit)
	// returns artists
	req, err := http.NewRequest("GET", URL+reqType, nil)
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", User.TokenType+" "+User.AccessToken)

	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading body. ", err)
	}
	json.Unmarshal(body, &list)
}
