package botmultiplexer

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// Translator component to handle translation of a HTTP payload into a
// consistent format, and to translate that format back into a HTTP payload
// for posting

func TranslateToProps(req *http.Request, propType string) (SessionData, bool) {

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("Error occurred reading http request: %v", err)
		return SessionData{}, false
	}
	log.Printf("Request body: %s", strings.ReplaceAll(string(reqBody), "\n", "\t"))

	switch propType {
	case TYPE_TELEGRAM:
		return TelegramTranslate(reqBody), true
	}

	return SessionData{}, true
}

func PostFromProps(env SessionData) bool {
	switch env.Type {
	case TYPE_TELEGRAM:
		log.Printf("Posting to Telegram -> {Message: %s, Affordances: %v}", env.Res.Message, env.Res.Affordances)
		return PostTelegram(env)
	}

	return false
}
