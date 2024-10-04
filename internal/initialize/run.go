package initialize

import (
	"fmt"
	"github.com/oeservices/open-kafka-connector/global"
	"go.uber.org/zap"
)

func Run() {
	LoadEnv()
	InitLogger()
	InitRouter()
	SetupDB()

	r := InitRouter()
	err := r.Run(fmt.Sprintf(":%d", global.Config.Server.Port))
	if err != nil {
		global.Logger.Error("Server error", zap.Error(err))
		panic(err)
		return
	}
}
