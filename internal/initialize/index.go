package initialize

import (
	"HoneyOrangeHelper/internal/config"
)

func Init() {
	InitRuntime()
	config.InitConfig()
	InitLogger()
	InitEvetnBus()
	InitPlugin()
}
