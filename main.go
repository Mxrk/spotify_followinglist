package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

var User access
var list List
var auth string

func main() {
	testTemplate, err := template.ParseFiles("list.gohtml")
	if err != nil {
		panic(err)
	}

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.HandleFunc("/callback", completeAuth)

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		log.Println("Got request for:", r.URL.String())
		id := strings.TrimSpace(r.FormValue("ClientID"))
		secret := strings.TrimSpace(r.FormValue("ClientSecret"))
		if id != "" && secret != "" {
			s := ReqPerms(id, secret)
			http.Redirect(w, r, s, http.StatusTemporaryRedirect)
		} else {
			// add error
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		}

	})

	http.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		err := testTemplate.Execute(w, list.Artists.Items)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.ListenAndServe(":8080", nil)

}
func completeAuth(w http.ResponseWriter, r *http.Request) {
	auth = r.FormValue("code")
	getAccessToken()
	http.Redirect(w, r, "/list", http.StatusTemporaryRedirect)
}
