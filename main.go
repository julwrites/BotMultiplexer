package BotMultiplexer

// Translator methods
import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	secrets := LoadSecrets()

	http.HandleFunc("/"+secrets.TELEGRAM_ID, ScriptureBot.telegramHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
