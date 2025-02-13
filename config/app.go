package config

import (
	"log"

	"github.com/spf13/viper"
)

// Env has environment stored
type Env struct {
	APP_NAME  string `mapstructure:"APP_NAME"`
	APP_DEBUG bool   `mapstructure:"APP_DEBUG"`
	APP_ENV   string `mapstructure:"APP_ENV"`

	SERVER_PORT string `mapstructure:"SERVER_PORT"`
	DBUrl       string `mapstructure:"DATABASE_URL"`
	TimeZone    string `mapstructure:"TZ"`
	LogStack     string `mapstructure:"LOG_STACK"`
	LogRetention int    `mapstructure:"LOG_RETENTION"`
}

// NewEnv creates a new environment
func NewEnv() Env {
	env := Env{}
	viper.SetConfigName(".env") // name of config file (without extension)
	viper.SetConfigType("env")  // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")    // optionally look for config in the working directory
	viper.AutomaticEnv()        // read in environment variables that match

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalf("☠️ .env file not found: %+v", err)
		} else {
			log.Fatalf("☠️ Error reading .env file: %+v", err)
		}
	}

	if err := viper.Unmarshal(&env); err != nil {
		log.Fatalf("☠️ environment can't be loaded: %+v", err)
	}

	if env.TimeZone == "" {
		env.TimeZone = "UTC"
	}

	return env
}
