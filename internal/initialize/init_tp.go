package initialize

import (
	"HoneyOrangeHelper/internal/global"
	"HoneyOrangeHelper/internal/plugin"
)

func InitPlugin() {
	srv, err := plugin.NewService()
	if err != nil {
		panic(err)
	}

	global.PluginSrv = srv
}
