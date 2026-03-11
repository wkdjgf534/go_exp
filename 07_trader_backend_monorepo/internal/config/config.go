package config

import (
	"fmt"

	"github.com/kkyr/fig"
)

type Config struct {
	MongoConfig MongoConfig  `fig:"mongo" validate:"required"`
	HTTPConfig  ServerConfig `fig:"http" validate:"required"`
}

type MongoConfig struct {
	Host        string `fig:"host" validate:"required"`
	Port        string `fig:"port" validate:"required"`
	Username    string `fig:"username"`
	Password    string `fig:"password"`
	Params      string `fig:"params"`
	Database    string `fig:"database" validate:"required"`
	AppName     string `fig:"appName" validate:"required"`
	MinPoolSize int    `fig:"minPoolSize" validate:"required"`
	MaxPoolSize int    `fig:"maxPoolSize" validate:"required"`
}

func (c MongoConfig) GetConnectionURI() string {
	result := fmt.Sprintf("mongodb://%s:%s", c.Host, c.Port)

	if c.Username != "" && c.Password != "" {
		result = fmt.Sprintf("mongodb://%s:%s@%s:%s", c.Username, c.Password, c.Host, c.Port)
	}

	if c.Params != "" {
		result += fmt.Sprintf("/%s", c.Params)
	}

	return result
}

type ServerConfig struct {
	Port int `fig:"port"`
}

func GetConfig() (*Config, error) {
	var config Config
	if err := fig.Load(&config, fig.UseEnv("")); err != nil {
		return nil, fmt.Errorf("error loading configuration: %w", err)
	}

	return &config, nil
}
