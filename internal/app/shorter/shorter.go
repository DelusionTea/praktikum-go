package shorter

import (
	"crypto/sha1"
	"encoding/base64"
	"strings"
)

func Shorter(longURL string) string {
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

//func ShorterGet(id int) string {
//	return fmt.Sprintf("%s%d", conf.ServerAddress, id)
//}
