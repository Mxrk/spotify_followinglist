package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

var cli string
var sec string

type access struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

type List struct {
	Artists struct {
		Items []struct {
			Followers struct {
				Href  interface{} `json:"href"`
				Total int         `json:"total"`
			} `json:"followers"`
			Genres []string `json:"genres"`
			Href   string   `json:"href"`
			ID     string   `json:"id"`
			Images []struct {
				Height int    `json:"height"`
				URL    string `json:"url"`
				Width  int    `json:"width"`
			} `json:"images"`
			Name       string `json:"name"`
			Popularity int    `json:"popularity"`
			Type       string `json:"type"`
			URI        string `json:"uri"`
		} `json:"items"`
	}
}

//Authorization Code Flow
//ReqPerms authorizes the app
func ReqPerms(clientID, clientSecret string) string {
	cli = clientID
	sec = clientSecret
	// need user-follow-read
	scope := "&scope=user-follow-read"
	url := "https://accounts.spotify.com/authorize?"
	responseType := "&response_type=code"
	redirectURI := "&redirect_uri=http://localhost:8080/callback"
	id := "&client_id=" + cli
	reqURL := url + scope + responseType + redirectURI + id
	return reqURL
}

func getAccessToken() {
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", auth)
	data.Set("redirect_uri", "http://localhost:8080/callback")
	data.Set("client_id", cli)
	data.Set("client_secret", sec)

	fmt.Println("Request with the following token: " + auth)

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))
	if err != nil {
		log.Println(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	json.Unmarshal(body, &User)
	getArtists()

}
