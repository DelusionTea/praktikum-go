package memory

import (
	"bufio"
	"encoding/json"
	"github.com/DelusionTea/praktikum-go/cmd/conf"
	"log"
	"os"
)

type MemoryMap struct {
	Values   map[string]string
	FilePath string
}

type row struct {
	ShortURL string `json:"short_url"`
	LongURL  string `json:"long_url"`
}

func (repo *MemoryMap) WriteRow(longURL string, shortURL string, filePath string) error {
	log.Println("Start write Row")
	file, err := os.OpenFile(repo.FilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, conf.FilePerm)

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
	if !reader.Scan() {
		log.Println("error read row Scan:  ")
		log.Print(reader.Err())
		return false, reader.Err()
	}
	reader.Scan()
	data := reader.Bytes()

	row := &row{}

	err := json.Unmarshal(data, row)

	if err != nil {
		log.Println("error read row  Unmarshal:  ")
		log.Print(err)
		return false, err
	}
	repo.Values[row.ShortURL] = row.LongURL
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
		Values:   values,
		FilePath: filePath,
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

func NewMemoryFile(filePath string) MemoryMap {
	log.Println("New Memory Map: ")
	//log.Print(NewMemoryMap(filePath))
	return *NewMemoryMap(filePath)
}
