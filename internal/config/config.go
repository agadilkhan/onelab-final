package config

import (
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Token TokenConfig `yaml:"token"`
	HTTP ServerConfig `yaml:"http"`
	DB   DBConfig     `yaml:"db"`
}

type TokenConfig struct {
	SecretKey string `env:"SECRET_KEY" env-required:"true"`
	TTL time.Duration `yaml:"ttl"`
}

type ServerConfig struct {
	Port            string `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout"`
	ReadTimeout time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
}

type DBConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Username string `yaml:"username"`
	DBName string `yaml:"db_name"`
	Password string `env:"DB_PASSWORD" env-required:"true"`
}

func InitConfig(path string) (*Config, error) {
	cfg := new(Config)
	
	err := cleanenv.ReadConfig(path, cfg)
	if err != nil {
		return nil, err
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}