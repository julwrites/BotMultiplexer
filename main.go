package main

// Translator methods
import (
	"fmt"
	"log"
	"net/http"
	"os"

	botsecrets "github.com/julwrites/BotSecrets"
	scripturebot "github.com/julwrites/ScriptureBot"
)

func multiplexer(res http.ResponseWriter, req *http.Request) {
	secrets := botsecrets.LoadSecrets()

	log.Printf("Url: %s", req.URL.EscapedPath())

	if req.URL.EscapedPath() == "/"+secrets.TELEGRAM_ID {
		log.Printf("Telegram message")

		scripturebot.TelegramHandler(res, req, &secrets)

		return
	}
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
