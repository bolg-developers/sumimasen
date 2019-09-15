package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
)

var env Environment

type Environment struct {
	ServerPort string `split_words:"true",required:"true"`
}

func Env() *Environment {
	return &env
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("dotenv error:", err)
	}
	if err := envconfig.Process("", &env); err != nil {
		log.Fatal("envconfig error:", err)
	}
}