package initialize

import (
	"github.com/joho/godotenv"
	"github.com/oeservices/open-kafka-connector/global"
	"github.com/oeservices/open-kafka-connector/pkg/settings"
	"os"
	"strconv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	serverPort, _ := strconv.Atoi(os.Getenv("SERVER_PORT"))
	global.Config = &settings.Config{
		Server: settings.ServerSettings{
			Port: serverPort,
			Mode: os.Getenv("SERVER_MODE"),
		},
		Kafka: settings.KafkaSettings{
			Broker: os.Getenv("KAFKA_BROKER"),
		},
	}
}
