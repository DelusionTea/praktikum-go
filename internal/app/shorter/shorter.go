package shorter

import (
	"crypto/sha1"
	"encoding/base64"
	"log"
	"strings"
)

func Shorter(longURL string) string {
	log.Println("Start Shorter")
	splitURL := strings.Split(longURL, "://")
	hashURL := sha1.New()
	if len(splitURL) < 2 {
		hashURL.Write([]byte(longURL))
	} else {
		hashURL.Write([]byte(splitURL[1]))
	}
	urlHash := base64.URLEncoding.EncodeToString(hashURL.Sum(nil))
	return string(urlHash)
}
