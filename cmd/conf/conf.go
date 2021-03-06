package conf

import (
	"flag"
	"github.com/caarlos0/env"
	"log"
	"os"
)

const (
	FilePerm    = 0755
	NumbWorkers = 10
	WorkerBuff  = 100
)

type Config struct {
	ServerAddress string `env:"SERVER_ADDRESS" envDefault:"localhost:8080"`
	BaseURL       string `env:"BASE_URL" envDefault:"http://localhost:8080/"`
	FilePath      string `env:"FILE_STORAGE_PATH" envDefault:"sorter.log"`
	DataBase      string `env:"DATABASE_DSN" envDefault:"" `
	Key           []byte
	NumbWorkers   int `env:"NUMBER_OF_WORKERS"`
	WorkerBuff    int `env:"WORKERS_BUFFER"`
}

func GetConfig() *Config {
	log.Println("Start Get Config")
	instance := &Config{}
	if err := env.Parse(instance); err != nil {
		log.Fatal(err)
	}
	ServerAddress := flag.String("a", instance.ServerAddress, "Server address")
	BaseURL := flag.String("b", instance.BaseURL, "base url")
	FileName := flag.String("f", instance.FilePath, "file path")

	DataBase := flag.String("d", instance.DataBase, "DataBase")
	flag.Parse()

	if os.Getenv("SERVER_ADDRESS") == "" {
		instance.ServerAddress = *ServerAddress
	}
	if os.Getenv("BASE_URL") == "" {
		instance.BaseURL = *BaseURL
	}
	if os.Getenv("FILE_STORAGE_PATH") == "" {
		instance.FilePath = *FileName
	}
	if os.Getenv("DATABASE_DSN") == "" {
		instance.DataBase = *DataBase
	}

	instance.Key = make([]byte, 16)
	instance.NumbWorkers = NumbWorkers
	instance.WorkerBuff = WorkerBuff
	log.Flags()
	log.Println(instance.BaseURL)
	log.Println(instance.ServerAddress)
	log.Println(instance.FilePath)

	if len(instance.BaseURL) > 0 {
		if instance.BaseURL[len(instance.BaseURL)-1] != '/' {
			instance.BaseURL += "/"
		}
	}
	return instance
}
