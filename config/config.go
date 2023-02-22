package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	PORT         string `json:"port"`
	JWTSecret    string `json:"jwt_secret"`
	MongoURI     string `json:"mongoURI"`
	ExchangeRate string `json:"exchange_rate"`
}

// Config contains configuration for console web server.

// InitDBConfigs gets environment variables needed to run the app
func InitDBConfigs() Config {

	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")
	jwtSecret := os.Getenv("JWT_SECRET")
	mongoURI := os.Getenv("MONGO_URI")
	exchangeRate := os.Getenv("EXCHANGE_RATE")

	return Config{
		PORT:         port,
		JWTSecret:    jwtSecret,
		MongoURI:     mongoURI,
		ExchangeRate: exchangeRate,
	}
}
