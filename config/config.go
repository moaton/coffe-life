package config

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		Postgres `yaml:"postgres"`
		HTTP     `yaml:"http"`
		JwtToken `yaml:"jwt_token"`
	}
	Postgres struct {
		Host         string `env-required:"true" yaml:"postgres_host" env:"POSTGRES_HOST"`
		Port         uint16 `env-required:"true" yaml:"postgres_port" env:"POSTGRES_PORT"`
		User         string `env-required:"true" yaml:"postgres_user" env:"POSTGRES_USER"`
		Password     string `env-required:"true" yaml:"postgres_password" env:"POSTGRES_PASSWORD"`
		DBName       string `env-required:"true" yaml:"postgres_db_name" env:"POSTGRES_DB_NAME"`
		MaxIdleConns int    `env-required:"true" yaml:"postgres_max_idle_conns" env:"POSTGRES_MAX_IDLE_CONNS"`
		MaxOpenConns int    `env-required:"true" yaml:"postgres_max_open_conns" env:"POSTGRES_MAX_OPEN_CONNS"`
	}
	HTTP struct {
		Server struct {
			Port string `env-required:"true" yaml:"http_server_port" env:"HTTP_SERVER_PORT"`
		} `yaml:"server"`
	}

	JwtToken struct {
		JwtTokenSecret string        `env-required:"true" yaml:"jwt_token_secret" env:"JWT_TOKEN_SECRET"`
		TokenTimeLimit time.Duration `env-required:"true" yaml:"jwt_token_time_limit" env:"JWT_TOKEN_TIME_LIMIT"`
	}
)

func (p Postgres) DSN() string {
	return fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", url.QueryEscape(p.User), url.QueryEscape(p.Password), p.Host, p.DBName)
}

func (j JwtToken) GetJwtTokenSecret() string {
	secret := j.JwtTokenSecret

	return secret
}

func (j JwtToken) GetTokenTimeLimit() time.Duration {
	return j.TokenTimeLimit
}

func New() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yaml", cfg)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("failed to load env from file: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
