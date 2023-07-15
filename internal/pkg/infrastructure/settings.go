package infrastructure

import (
	"fmt"
	"log"
	"os"
)

type Env struct {
	PORT               string
	DOMAIN             string
	REPOSITORY_ADAPTER string
}

var env *Env

func Settings() *Env {
	if env != nil {
		return env
	}
	env = &Env{
		PORT:               os.Getenv("PORT"),
		DOMAIN:             os.Getenv("DOMAIN"),
		REPOSITORY_ADAPTER: os.Getenv("REPOSITORY_ADAPTER"),
	}
	if env.PORT == "" {
		log.Println("Solving port...")
		env.PORT = "3333"
	}
	if env.DOMAIN == "" {
		log.Println("Solving domain...")
		env.DOMAIN = fmt.Sprintf("http://localhost:%s", env.PORT)
	}
	if env.REPOSITORY_ADAPTER == "" {
		log.Println("Solving repository adapter...")
		env.REPOSITORY_ADAPTER = "in_memory"
	}
	log.Println("Environment variables setted")
	return env
}
