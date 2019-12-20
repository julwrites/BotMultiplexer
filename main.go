package main

// Translator methods
import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	botsecrets "github.com/julwrites/BotSecrets"
	scripturebot "github.com/julwrites/ScriptureBot"
)

func multiplexer(res http.ResponseWriter, req *http.Request) {
	secretsPath, _ := filepath.Abs("./secrets.yaml")
	secrets := botsecrets.LoadSecrets(secretsPath)

	log.Printf("URL: %s", req.URL.EscapedPath())

	log.Printf("Telegram: %s", "/"+secrets.TELEGRAM_ID)
	if strings.Compare(strings.Trim(req.URL.EscapedPath(), "\n"), strings.Trim("/"+secrets.TELEGRAM_ID, "\n")) == 0 {
		log.Printf("Incoming telegram message")

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
