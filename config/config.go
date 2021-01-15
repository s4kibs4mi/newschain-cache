package config

import (
	"github.com/spf13/viper"
	"sync"
)

var mu sync.Mutex

// LoadConfig initiates of config load
func LoadConfig() error {
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	LoadApp()

	return nil
}
