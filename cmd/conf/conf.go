package conf

import (
	"flag"
	"github.com/caarlos0/env"
	"log"
	"os"
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
	if err := env.Parse(instance); err != nil {
		log.Fatal(err)
	}
	//conf.BaseURL = fmt.Sprintf("http://%s/", conf.ServerAddress)
	//if err := env.Parse(&conf); err != nil {
	//	log.Fatal(err)
	//}

	//ServerAddress := flag.String(instance.ServerAddress, "a", "Server address")
	//BaseURL := flag.String(instance.BaseURL, "b", "base url")
	//FileName := flag.String(instance.FilePath, "f", "file path")
	ServerAddress := flag.String("a", instance.ServerAddress, "Server address")
	BaseURL := flag.String("b", instance.BaseURL, "base url")
	FileName := flag.String("f", instance.FilePath, "file path")
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

	log.Flags()
	log.Println(instance.BaseURL)
	log.Println(instance.ServerAddress)
	log.Println(instance.FilePath)
	return instance
}
