package memory

import (
	"bufio"
	"encoding/json"
	"errors"
	"github.com/DelusionTea/praktikum-go/cmd/conf"
	"github.com/DelusionTea/praktikum-go/internal/app/Shorter"
	"log"
	"os"
)

type MemoryInterface interface {
	AddURL(longURL string) string
	GetURL(shortURL string) (string, error)
}
type MemoryMap struct {
	values   map[string]string
	filePath string
}

type row struct {
	ShortURL string `json:"short_url"`
	LongURL  string `json:"long_url"`
}

func (repo *MemoryMap) writeRow(longURL string, shortURL string, filePath string) error {
	file, err := os.OpenFile(repo.filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, conf.FilePerm)

	if err != nil {
		return err
	}
	writer := bufio.NewWriter(file)

	data, err := json.Marshal(&row{
		LongURL:  longURL,
		ShortURL: shortURL,
	})
	if err != nil {
		return err
	}

	if _, err := writer.Write(data); err != nil {
		return err
	}

	if err := writer.WriteByte('\n'); err != nil {
		return err
	}

	return writer.Flush()
}

func (repo *MemoryMap) readRow(reader *bufio.Scanner) (bool, error) {

	if !reader.Scan() {
		return false, reader.Err()
	}
	data := reader.Bytes()

	row := &row{}

	err := json.Unmarshal(data, row)

	if err != nil {
		return false, err
	}
	repo.values[row.ShortURL] = row.LongURL

	return true, nil
}

func NewMemoryMap(filePath string) *MemoryMap {
	values := make(map[string]string)
	repo := MemoryMap{
		values:   values,
		filePath: filePath,
	}
	file, err := os.OpenFile(filePath, os.O_RDONLY|os.O_CREATE, conf.FilePerm)
	if err != nil {
		log.Printf("Error with reading file: %v\n", err)
	}
	defer file.Close()
	reader := bufio.NewScanner(file)

	for {
		ok, err := repo.readRow(reader)

		if err != nil {
			log.Printf("Error while parsing file: %v\n", err)
		}

		if !ok {
			break
		}
	}

	return &repo
}

func (repo *MemoryMap) AddURL(longURL string) string {
	shortURL := Shorter.Shorter(longURL)
	repo.values[shortURL] = longURL
	repo.writeRow(longURL, shortURL, repo.filePath)
	return shortURL
}

func (repo *MemoryMap) GetURL(shortURL string) (string, error) {
	resultURL, okey := repo.values[shortURL]
	if !okey {
		return "", errors.New("not found")
	}
	return resultURL, nil
}

func NewMemoryFile(filePath string) MemoryInterface {
	return MemoryInterface(NewMemoryMap(filePath))
}
