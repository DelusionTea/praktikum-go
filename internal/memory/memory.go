package memory

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"github.com/DelusionTea/praktikum-go/cmd/conf"
	"github.com/DelusionTea/praktikum-go/internal/app/handlers"
	"log"
	"os"
)

type MemoryMap struct {
	Values   map[string]string
	FilePath string
	BaseURL  string
	UsersURL map[string][]string
}

type row struct {
	ShortURL string `json:"short_url"`
	LongURL  string `json:"long_url"`
	User     string `json:"user"`
}

func (repo *MemoryMap) WriteRow(longURL string, shortURL string, filePath string, user string) error {
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
		User:     user,
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

func NewMemoryMap(ctx context.Context, filePath string, baseURL string) *MemoryMap {
	repo := MemoryMap{
		Values:   map[string]string{},
		FilePath: filePath,
		BaseURL:  baseURL,
		UsersURL: map[string][]string{},
	}
	file, err := os.OpenFile(repo.FilePath, os.O_RDONLY|os.O_CREATE, conf.FilePerm)
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

func NewMemoryFile(ctx context.Context, filePath string, baseURL string) handlers.ShorterInterface {
	return handlers.ShorterInterface(NewMemoryMap(ctx, filePath, baseURL))
}
func (repo *MemoryMap) AddManyURL(ctx context.Context, urls []handlers.ManyPostURL, user string) ([]handlers.ManyPostResponse, error) {
	return nil, nil
}

func (repo *MemoryMap) AddURL(ctx context.Context, longURL string, shortURL string, user string) error {
	repo.Values[shortURL] = longURL
	repo.WriteRow(longURL, shortURL, repo.FilePath, user)
	repo.UsersURL[user] = append(repo.UsersURL[user], shortURL)
	return nil
}

func (repo *MemoryMap) GetURL(ctx context.Context, shortURL string) (string, error) {
	resultURL, okey := repo.Values[shortURL]
	if !okey {
		return "", errors.New("not found")
	}
	return resultURL, nil
}

func (repo *MemoryMap) GetUserURL(ctx context.Context, user string) ([]handlers.ResponseGetURL, error) {
	result := []handlers.ResponseGetURL{}
	for _, url := range repo.UsersURL[user] {
		temp := handlers.ResponseGetURL{
			ShortURL:    repo.BaseURL + url,
			OriginalURL: repo.Values[url],
		}
		result = append(result, temp)
	}

	return result, nil
}

func (repo *MemoryMap) Ping(ctx context.Context) error {
	return nil
}
