package memory

import (
	"bufio"
	"encoding/json"
	"errors"
	"github.com/DelusionTea/praktikum-go/cmd/conf"
	"github.com/DelusionTea/praktikum-go/internal/app/shorter"
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
	log.Println("Start write Row")
	file, err := os.OpenFile(repo.filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, conf.FilePerm)

	if err != nil {
		log.Println("error write row")
		return err
	}
	writer := bufio.NewWriter(file)

	data, err := json.Marshal(&row{
		LongURL:  longURL,
		ShortURL: shortURL,
	})
	if err != nil {
		log.Println("error write row")
		return err
	}

	if _, err := writer.Write(data); err != nil {
		log.Println("error write row")
		return err
	}

	if err := writer.WriteByte('\n'); err != nil {
		log.Println("error write row")
		return err
	}

	return writer.Flush()
}

func (repo *MemoryMap) readRow(reader *bufio.Scanner) (bool, error) {
	log.Println("Start read Row")
	//if !reader.Scan() {
	//	log.Println("error read row Scan:  ")
	//	log.Print(reader.Err())
	//	return false, reader.Err()
	//}
	reader.Scan()
	data := reader.Bytes()

	row := &row{}

	err := json.Unmarshal(data, row)

	if err != nil {
		log.Println("error read row  Unmarshal:  ")
		log.Print(err)
		return false, err
	}
	repo.values[row.ShortURL] = row.LongURL
	log.Println("readRow long URL:  ")
	log.Print(row.LongURL)

	log.Println("readRow ShortURL:  ")
	log.Print(row.ShortURL)
	return true, nil
}

func NewMemoryMap(filePath string) *MemoryMap {
	log.Println("Start Create Memory Map")
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
	log.Println("result of ReadRow: ")
	log.Print(&repo)
	return &repo
}

func (repo *MemoryMap) AddURL(longURL string) string {
	log.Println("Start Add URL")
	shortURL := shorter.Shorter(longURL)
	repo.values[shortURL] = longURL
	repo.writeRow(longURL, shortURL, repo.filePath)
	log.Println("End Add URL :")
	log.Print(shortURL)
	return shortURL
}

func (repo *MemoryMap) GetURL(shortURL string) (string, error) {
	log.Println("Start Get URL")
	resultURL, okey := repo.values[shortURL]
	log.Println("End Get URL :")
	log.Print(resultURL)
	if !okey {
		return "", errors.New("not found")
	}
	return resultURL, nil
}

func NewMemoryFile(filePath string) MemoryInterface {
	log.Println("New Memory Map: ")
	log.Print(MemoryInterface(NewMemoryMap(filePath)))
	return MemoryInterface(NewMemoryMap(filePath))
}
