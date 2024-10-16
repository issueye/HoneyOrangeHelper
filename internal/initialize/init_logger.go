package initialize

import (
	"HoneyOrangeHelper/internal/config"
	"HoneyOrangeHelper/internal/global"
	"HoneyOrangeHelper/pkg/logger"
	"fmt"
)

func InitLogger() {
	cfg := new(logger.Config)
	cfg.Level = config.GetParam(config.LOG, "level", 0).Int()
	cfg.Path = config.GetParam("log", "path", fmt.Sprintf("%s/logs", global.ROOT_PATH)).String()
	cfg.MaxSize = config.GetParam(config.LOG, "max-size", 100).Int() // MB
	cfg.MaxBackups = config.GetParam(config.LOG, "max-backups", 100).Int()
	cfg.MaxAge = config.GetParam(config.LOG, "max-age", 100).Int() // days
	cfg.Compress = config.GetParam(config.LOG, "compress", true).Bool()
	global.Sugared, global.Logger = logger.InitLogger(cfg)
}
