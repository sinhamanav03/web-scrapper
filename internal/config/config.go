package config

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jinzhu/configor"
	"github.com/joho/godotenv"
	"github.com/sinhamanav03/web-scrapper/internal/db"
)

var ErrInvalidFileExtension = errors.New("file extension not supported")

type Config struct {
	Server struct {
		Port int
		Host string
	}

	Database *db.Config
}

func Load(fileNames ...string) (*Config, error) {
	loadFiles := make([]string, 0, len(fileNames))
	envFiles := make([]string, 0, len(fileNames))

	for _, file := range fileNames {
		fileParts := strings.Split(file, ".")

		ext := fileParts[len(fileParts)-1]
		switch ext {
		case "yml", "json", "yaml", "toml":
			loadFiles = append(loadFiles, file)
		case "env":
			envFiles = append(envFiles, file)
		default:
			return nil, ErrInvalidFileExtension
		}
	}

	if len(envFiles) > 0 {
		err := godotenv.Load(envFiles...)
		if err != nil {
			return nil, fmt.Errorf("error while loading env files(%s): %w", strings.Join(envFiles, ","), err)
		}
	}

	cfg, err := loadConfig(loadFiles...)

	if err != nil {
		return nil, err
	}

	return cfg, err
}

func loadConfig(fileNames ...string) (*Config, error) {
	var config Config

	cfg := newConf()

	err := cfg.Load(&config, fileNames...)
	if err != nil {
		return nil, fmt.Errorf("cannot load config files(%s): %w", strings.Join(fileNames, ","), err)
	}

	return &config, nil
}

func newConf() *configor.Configor {
	conf := configor.Config{ENVPrefix: "GAUTH"}
	config := configor.New(&conf)

	return config
}
