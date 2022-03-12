package conf

import (
	"flag"
	"os"
)

const (
	ServerAddress = ":8080"
	BaseURL       = "http://localhost:8080/"
	FileName      = "sorter.logs"
	//FilePerm      = 0755
)

type Config struct {
	BaseURL       string
	ServerAddress string
	FilePath      string
}

func GetConfig() *Config {
	cfg := &Config{}
	if cfg.BaseURL = os.Getenv("SERVER_ADDRESS"); cfg.BaseURL == "" {
		flag.StringVar(&cfg.BaseURL, "a", ServerAddress, "Server address")
	}

	if cfg.ServerAddress = os.Getenv("BASE_URL"); cfg.ServerAddress == "" {
		flag.StringVar(&cfg.ServerAddress, "b", BaseURL, "Server base URL")
	}

	if cfg.FilePath = os.Getenv("FILE_STORAGE_PATH"); cfg.FilePath == "" {
		flag.StringVar(&cfg.FilePath, "f", FileName, "File storage path")
	}

	flag.Parse()

	return cfg
}
