package utils

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/subosito/gotenv"
)

type Config struct {
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBName   string `mapstructure:"DB_NAME"`
	DBHost   string `mapstructure:"DB_HOST"`
	DBPort   string `mapstructure:"DB_PORT"`
	DBUser   string `mapstructure:"DB_USER"`
	DBPass   string `mapstructure:"DB_PASS"`
	Address  string `mapstructure:"ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
	p := filepath.Join(path, ".env")
	err = gotenv.Load(p)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("Will Load from env and not file")
			err = nil
		} else {
			return
		}
	}
	config.DBDriver = os.Getenv("DB_DRIVER")
	config.DBName = os.Getenv("DB_NAME")
	config.DBHost = os.Getenv("DB_HOST")
	config.DBPort = os.Getenv("DB_PORT")
	config.DBUser = os.Getenv("DB_USER")
	config.DBPass = os.Getenv("DB_PASS")
	config.Address = os.Getenv("ADDRESS")
	return
}

func (config *Config) GetDBString() string {
	// return "postgresql://root:secret@localhost:5432/sample_bank?sslmode=disable"
	return fmt.Sprintf("%s://%s:%s@%s:%s/%s/sslmode=disable", config.DBDriver, config.DBUser, config.DBPass, config.DBHost, config.DBPort, config.DBName)
}
