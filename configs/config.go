package configs

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Port          uint          `yaml:"port"`
	DBConfig      DBConfig      `yaml:"db"`
	RedisConfig   RedisConfig   `yaml:"redis"`
	SessionConfig SessionConfig `yaml:"session"`
}

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Addr     string `yaml:"addr"`
	Port     int    `yaml:"port"`
	DBName   string `yaml:"dbname"`
	Password string `yaml:"password"`
}

type SessionConfig struct {
	ExpiresAt int `yaml:"expiresAt"`
}

const dbPasswordName = "DB_PASSWORD"

func InitConfig(config *Config) error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}

	config.DBConfig.Password = os.Getenv(dbPasswordName)

	filename, err := filepath.Abs("./configs/config.yaml")
	if err != nil {
		return err
	}

	yamlFile, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		return err
	}

	return nil
}
