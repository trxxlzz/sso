package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env          string        `yaml:"env" env-default:"local"`
	Storage_path string        `yaml:"storage_path" env-required:"true"`
	TokenTTl     time.Duration `yaml:"token_ttl" env-required:"true"`
	GRPC         GRPCConfig    `yaml:"grpc"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func MustLoad() *Config {
	ConfigPath := os.Getenv("CONFIG_PATH")
	if ConfigPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(ConfigPath); os.IsNotExist(err) {
		log.Fatal("CONFIG_PATH does not exist")
	}

	var cfg Config

	if err := cleanenv.ReadConfig(ConfigPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
