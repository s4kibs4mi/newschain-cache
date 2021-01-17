package config

import (
	"github.com/spf13/viper"
	"sync"
)

var mu sync.Mutex

// LoadConfig initiates of config load
func LoadConfig() error {
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	LoadApp()
	LoadDB()
	LoadEthereum()

	return nil
}
