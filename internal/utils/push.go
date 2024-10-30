package utils

import (
	"HoneyOrangeHelper/internal/config"
	"HoneyOrangeHelper/internal/global"
	"github.com/issueye/ipc_grpc/grpc/pb"
)

func PushEvent(t pb.EventType, info *pb.ServerInfo) {
	list, err := config.GetToolPluginList("")
	if err != nil {
		return
	}

	for _, v := range list {
		global.PluginSrv.Events.PushEvent(v.CookieKey, t, info)
	}
}
