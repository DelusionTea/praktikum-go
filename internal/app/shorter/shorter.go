package shorter

import (
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"github.com/DelusionTea/praktikum-go/internal/memory"
	"log"
	"strings"
)

type ShortnerInterface interface {
	AddURL(longURL string, repo memory.MemoryMap) string
	GetURL(shortURL string, repo memory.MemoryMap) (string, error)
}

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

func AddURL(longURL string, repo memory.MemoryMap) string {
	log.Println("Start Add URL")
	shortURL := Shorter(longURL)
	repo.Values[shortURL] = longURL
	repo.WriteRow(longURL, shortURL, repo.FilePath)
	log.Println("End Add URL :")
	log.Print(shortURL)
	return shortURL
}

func GetURL(shortURL string, repo memory.MemoryMap) (string, error) {
	log.Println("Start Get URL")
	resultURL, okey := repo.Values[shortURL]
	log.Println("End Get URL :")
	log.Print(resultURL)
	if !okey {
		return "", errors.New("not found")
	}
	return resultURL, nil
}
