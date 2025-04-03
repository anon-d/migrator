package config

import (
	"fmt"
	"log"

	env "github.com/caarlos0/env/v11"
	"github.com/go-playground/validator"
)

type Config struct {
	Driver             string `env:"GOOSE_DRIVER"`
	DBUrl              string `env:"GOOSE_DBSTRING"`
	MigrationDirectory string `env:"GOOSE_MIGRATION"`
}

func MustLoad() (*Config, error) {

	cfgPr := &Config{}

	if err := env.Parse(cfgPr); err != nil {
		log.Println("Error. Can't parse ENV")
		return cfgPr, fmt.Errorf("can't parse ENV: %w", err)
	}

	err := validator.New().Struct(cfgPr)
	if err != nil {
		log.Print("Validation error")
		return cfgPr, fmt.Errorf("validation error: %w", err)
	}

	return cfgPr, nil
}
