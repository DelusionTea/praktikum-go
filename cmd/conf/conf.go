package conf

import (
	"flag"
	"fmt"
	"github.com/caarlos0/env"
	"log"
)

const (
	//ServerAddress = "localhost:8080"
	//BaseURL       = "http://localhost:8080/"
	//FileName      = "sorter.logs"
	FilePerm = 0755
)

type Config struct {
	ServerAddress string `env:"SERVER_ADDRESS" envDefault:"localhost:8080"`
	BaseURL       string `env:"BASE_URL" envDefault:"http://localhost:8080/"`
	FilePath      string `env:"FILE_STORAGE_PATH" envDefault:"sorter.logs"`
}

var instance *Config

func GetConfig() *Config {
	log.Println("Start Get Config")
	instance = &Config{}
	instance.BaseURL = fmt.Sprintf("http://%s/", instance.ServerAddress)
	if err := env.Parse(&instance); err != nil {
		log.Fatal(err)
	}

	ServerAddress := flag.String(instance.ServerAddress, "a", "Server address")
	BaseURL := flag.String(instance.BaseURL, "b", "base url")

	FileName := flag.String(instance.FilePath, "f", "file path")
	flag.Parse()

	if instance.ServerAddress == "" {
		instance.ServerAddress = *ServerAddress
	}
	if instance.BaseURL == "" {
		instance.BaseURL = *BaseURL
	}
	if instance.FilePath == "" {
		instance.FilePath = *FileName
	}

	return instance
}
