package domain

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	AppName      string
	AppEnv       string
	Host         string
	Port         string
	Username     string
	Password     string
	Database     string
	JwtSecretKey []byte
}

var GlobalConfig AppConfig

func (cfg *AppConfig) LoadVariables(envPath ...string) error {
	err := godotenv.Load(envPath...)

	if err != nil {
		log.Println(".env file not found. Loading from system environment", err)
	}

	cfg.AppName = os.Getenv("APP_NAME")
	cfg.AppEnv = os.Getenv("APP_ENV")
	cfg.Host = os.Getenv("DB_HOST")
	cfg.Username = os.Getenv("DB_USERNAME")
	cfg.Password = os.Getenv("DB_PASSWORD")
	cfg.Database = os.Getenv("DB_DATABASE")
	cfg.Port = os.Getenv("DB_PORT")
	cfg.JwtSecretKey = []byte("mindawakebodyasleep")

	return nil
}

func getEnvInt(key string) int {
	val, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		log.Fatalln("Invalid key:", key, "It should be an integer")
	}

	return val
}
