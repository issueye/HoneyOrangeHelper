package config

import (
	"HoneyOrangeHelper/internal/global"
	"fmt"
)

const (
	SERVER = "server"
	LOG    = "log"
)

func InitConfig() {
	SetParamExist(SERVER, "port", "端口号", 10070)
	SetParamExist(SERVER, "mode", `服务运行模式， debug \ release`, "debug")

	SetParamExist(LOG, "path", "日志存放路径", fmt.Sprintf("%s/logs", global.ROOT_PATH))
	SetParamExist(LOG, "max-size", "日志大小", 10)
	SetParamExist(LOG, "max-backups", "最大备份数", 10)
	SetParamExist(LOG, "max-age", "保存天数", 10)
	SetParamExist(LOG, "compress", "是否压缩", true)
	SetParamExist(LOG, "level", "日志输出等级", -1)
}
