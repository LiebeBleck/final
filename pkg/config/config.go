package config

import "os"

var Config struct {
	Password string
}

func Init() {
	Config.Password = os.Getenv("TODO_PASSWORD")
}
