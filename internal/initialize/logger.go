package initialize

import (
	"github.com/oeservices/open-kafka-connector/global"
	"github.com/oeservices/open-kafka-connector/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger("debug")
}
