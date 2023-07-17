package config

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

type Settings struct {
	ENV                string
	PORT               string
	DOMAIN             string
	REPOSITORY_ADAPTER string
	MEMCACHED_URL      string
	REDIS_URL          string
	REDIS_PASSWORD     string
	REDIS_DB           int
}

var settings *Settings

func GetSettings() *Settings {
	if settings != nil {
		return settings
	}

	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Join(filepath.Dir(b), "../../../..")

	viper.SetConfigName("config")
	viper.AddConfigPath(basepath + "/configs/")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	viper.AutomaticEnv()
	viper.SetDefault("PORT", "3333")
	viper.SetDefault("REPOSITORY_ADAPTER", "in_memory")

	settings = &Settings{
		ENV:                viper.GetString("ENV"),
		PORT:               viper.GetString("PORT"),
		DOMAIN:             viper.GetString("DOMAIN"),
		REPOSITORY_ADAPTER: viper.GetString("REPOSITORY_ADAPTER"),
		MEMCACHED_URL:      viper.GetString("MEMCACHED_URL"),
		REDIS_URL:          viper.GetString("REDIS_URL"),
		REDIS_PASSWORD:     viper.GetString("REDIS_PASSWORD"),
		REDIS_DB:           viper.GetInt("REDIS_DB"),
	}

	if settings.DOMAIN == "" {
		log.Println("Solving domain...")
		settings.DOMAIN = fmt.Sprintf("http://localhost:%s", settings.PORT)
	}

	log.Printf("Environment variables setted in %s environment", settings.ENV)
	return settings
}
