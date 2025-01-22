package config

import (
	"github.com/spf13/viper"
	"log"
)

// Env has environment stored
type Env struct {
	APP_NAME  string `mapstructure:"APP_NAME"`
	APP_DEBUG bool   `mapstructure:"APP_DEBUG"`

	SERVER_PORT string `mapstructure:"SERVER_PORT"`
	DBUrl       string `mapstructure:"DATABASE_URL"`
	TimeZone    string `mapstructure:"TZ"`
}

type EnvPath string

func (p EnvPath) ToString() string {
	return string(p)
}

// NewEnv creates a new environment
func NewEnv(envPath EnvPath) Env {
	env := Env{}
	viper.SetConfigFile(envPath.ToString())

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalf("☠️ Env config file not found: %+v", err)
		} else {
			log.Fatalf("☠️ Env config file error: %+v", err)
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
