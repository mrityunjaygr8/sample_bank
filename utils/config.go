package utils

import (
	"fmt"

	"github.com/spf13/viper"
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
	viper.AddConfigPath(path)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

func (config *Config) GetDBString() string {
	// return "postgresql://root:secret@localhost:5432/sample_bank?sslmode=disable"
	return fmt.Sprintf("%s://%s:%s@%s:%s/%s/sslmode=disable", config.DBDriver, config.DBUser, config.DBPass, config.DBHost, config.DBPort, config.DBName)
}
