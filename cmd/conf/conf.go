package conf

import (
	"flag"
	"fmt"
	"github.com/caarlos0/env"
	"log"
)

const (
	ServerAddress = ":8080"
	BaseURL       = "http://localhost:8080/"
	FileName      = "sorter.logs"
	FilePerm      = 0755
)

type Config struct {
	ServerAddress string `env:"SERVER_ADDRESS"`
	BaseURL       string `env:"BASE_URL"`
	FilePath      string `env:"FILE_STORAGE_PATH"`
}

func GetConfig() *Config {
	log.Println("Start Get Config")
	conf := Config{
		ServerAddress: ServerAddress,
		FilePath:      FileName,
		BaseURL:       BaseURL,
	}
	conf.BaseURL = fmt.Sprintf("http://%s/", conf.ServerAddress)
	if err := env.Parse(&conf); err != nil {
		log.Fatal(err)
	}

	flag.StringVar(&conf.ServerAddress, "a", ServerAddress, "Server address")
	flag.StringVar(&conf.BaseURL, "b", BaseURL, "base url")

	flag.StringVar(&conf.FilePath, "f", FileName, "file path")
	//flag.Parse()

	if conf.ServerAddress == "" {
		conf.ServerAddress = ServerAddress
	}
	if conf.BaseURL == "" {
		conf.BaseURL = BaseURL
	}
	if conf.FilePath == "" {
		conf.FilePath = FileName
	}

	if string(conf.BaseURL[len(conf.BaseURL)-1]) != "/" {
		conf.BaseURL += "/"
	}

	return &conf
}
