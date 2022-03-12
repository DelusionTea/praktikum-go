package conf

import (
	"flag"
)

const (
	ServerAddress = ":8080"
	BaseURL       = "http://localhost:8080/"
	FileName      = "sorter.logs"
	//FilePerm      = 0755
)

type Config struct {
	BaseURL       string `env:"BASE_URL"`
	ServerAddress string `env:"SERVER_ADDRESS"`
	FilePath      string `env:"FILE_STORAGE_PATH"`
}

func GetConfig() *Config {
	cfg := &Config{}
	if cfg.BaseURL == "" {
		flag.StringVar(&cfg.BaseURL, "a", ServerAddress, "Server address")
	}

	if cfg.ServerAddress == "" {
		flag.StringVar(&cfg.ServerAddress, "b", BaseURL, "Server base URL")
	}

	if cfg.FilePath == "" {
		flag.StringVar(&cfg.FilePath, "f", FileName, "File storage path")
	}

	flag.Parse()

	return cfg
}
