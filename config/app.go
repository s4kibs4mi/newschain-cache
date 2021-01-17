package config

import (
	"github.com/spf13/viper"
)

// Application holds the application configuration
type Application struct {
	Host string
	Port int
	Mode string
}

// app is the default application configuration
var app Application

// App returns the default application configuration
func App() *Application {
	return &app
}

// LoadApp loads application configuration
func LoadApp() {
	mu.Lock()
	defer mu.Unlock()

	app = Application{
		Host: viper.GetString("app.host"),
		Port: viper.GetInt("app.port"),
		Mode: viper.GetString("app.mode"),
	}
}
