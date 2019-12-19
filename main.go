package main

// Translator methods
import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func multiplexer(res http.ResponseWriter, req *http.Request) bool {
	secrets := LoadSecrets()

	if req.URL.Path == ("/" + secrets.TELEGRAM_ID) {
		log.Printf("Telegram message")

		ScriptureBot.telegramHandler(res, req, secrets)

		return true
	}

	return false
}

func main() {

	http.HandleFunc("/", multiplexer)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
