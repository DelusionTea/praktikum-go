package conf

import (
	"flag"
	"fmt"
	"github.com/caarlos0/env"
	"log"
	"os"
	"path/filepath"
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

	flagAddress := flag.String("a", ServerAddress, "Server address")
	flagBaseURL := flag.String("b", BaseURL, "base url")
	flagFilePath := flag.String("f", FileName, "file path")
	//flag.Parse()

	if *flagAddress != ServerAddress {
		conf.ServerAddress = *flagAddress
	}
	if *flagBaseURL != BaseURL {
		conf.BaseURL = *flagBaseURL
	}
	if *flagFilePath != FileName {
		conf.FilePath = *flagFilePath
	}

	if conf.FilePath != FileName {
		if _, err := os.Stat(filepath.Dir(conf.FilePath)); os.IsNotExist(err) {
			log.Println("Creating folder")
			err := os.Mkdir(filepath.Dir(conf.FilePath), FilePerm)
			if err != nil {
				log.Printf("Error: %v \n", err)
			}
		}
	}

	if string(conf.BaseURL[len(conf.BaseURL)-1]) != "/" {
		conf.BaseURL += "/"
	}

	return &conf
}
