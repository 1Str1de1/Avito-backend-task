package server

import (
	"fmt"
	"os"
	"reflect"
)

type Config struct {
	bindAddr    string `yaml:"bind_addr" bind_addr-default:"localhost:8080"`
	storagePath string `yaml:"storage_path"`
	env         *Env
}

// Env defines environment variables necessary for the app to run
type Env struct {
	ApiVersion       string `env:"API_VERSION"`
	ApiPort          string `env:"API_PORT"`
	PostgresUser     string `env:"POSTGRES_USER"`
	PostgresPassword string `env:"POSTGRES_PASSWORD"`
	PostgresName     string `env:"POSTGRES_NAME"`
	PostgresHost     string `env:"POSTGRES_HOST"`
}

// getEnv gets all necessary environment variables
func getEnv() *Env {
	env := Env{}
	v := reflect.ValueOf(&env).Elem()
	t := reflect.TypeOf(env)
	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag.Get("env")
		val := os.Getenv(tag)
		if len(val) == 0 {
			panic(fmt.Sprintf("Error getting environment variable: %s not set", tag))
		}
		v.Field(i).SetString(val)
	}
	return &env
}

func NewConfig() *Config {
	return &Config{
		bindAddr: ":8080",
		env:      getEnv(),
	}
}
