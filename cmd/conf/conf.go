package conf

import (
	"flag"
	"os"
)

type Config struct {
	BaseURL       string
	ServerAddress string
	FilePath      string
}

func GetConfig() *Config {
	cfg := &Config{}
	if cfg.BaseURL = os.Getenv("SERVER_ADDRESS"); cfg.BaseURL == "" {
		flag.StringVar(&cfg.BaseURL, "a", ":8080", "Server address")
	}

	if cfg.ServerAddress = os.Getenv("BASE_URL"); cfg.ServerAddress == "" {
		flag.StringVar(&cfg.ServerAddress, "b", "http://localhost:8080/", "Server base URL")
	}

	if cfg.FilePath = os.Getenv("FILE_STORAGE_PATH"); cfg.FilePath == "" {
		flag.StringVar(&cfg.FilePath, "f", "", "File storage path")
	}

	flag.Parse()

	return cfg
}
