package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/peti2001/flashSecretSlackBot/adapter"
)

type flashSecret interface {
	SaveSecret(string) (string, error)
}

func getHandler(fs flashSecret, token string) func(http.ResponseWriter, *http.Request) {
	response := `{ "text": "URL to share the secret: %s" }`
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		log.Printf("Request from user: %s\n", r.FormValue("user_name"))
		if token != "" && r.FormValue("token") != token {
			log.Println("Wrong Slack token")
			w.Write([]byte("Authentication problem. Wrong Slack token was provided."))
			return
		}
		url, err := fs.SaveSecret(r.FormValue("text"))
		if err != nil {
			fmt.Println(err)
			w.Write([]byte("There was an error. The service is not availablbe."))
			return

		}

		js := fmt.Sprintf(response, url)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(js))
	}
}

func main() {
	var fs flashSecret

	path := flag.String("path", "/", "The URL where the service will be availabble")
	port := flag.Int("port", 8080, "The port where the service will be available")
	slackToken := flag.String("slack-token", "", "A secret token from Slack. Leave it empty to skip the check.")
	flag.Parse()
	fs = adapter.QuickForgetFactory()

	http.HandleFunc(*path, getHandler(fs, *slackToken))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil); err != nil {
		fmt.Println(err)
	}
}
