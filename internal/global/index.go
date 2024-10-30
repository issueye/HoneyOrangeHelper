package global

import (
	"HoneyOrangeHelper/internal/plugin"

	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"go.uber.org/zap"
)

var (
	MsgChannel = make(chan string, 10)
	PubSub     *gochannel.GoChannel
	Sugared    *zap.SugaredLogger
	Logger     *zap.Logger
	PluginSrv  *plugin.Service
)

const (
	TOPIC_CONSOLE_LOG                = "TOPIC_CONSOLE_LOG"
	TOPIC_SERVER_REMOVE              = "TOPIC_SERVER_REMOVE"
	TOPIC_SERVER_INDEX               = "TOPIC_SERVER_INDEX"
	TOPIC_SERVER_REFRESH_TOOL_PLUGIN = "TOPIC_SERVER_REFRESH_TOOL_PLUGIN"
	TOPIC_SERVER_LOAD_TOOL_PLUGIN    = "TOPIC_SERVER_LOAD_TOOL_PLUGIN"
	TOPIC_SERVER_UNLOAD_TOOL_PLUGIN  = "TOPIC_SERVER_UNLOAD_TOOL_PLUGIN"
	TOPIC_SERVER_HELPER              = "TOPIC_SERVER_HELPER"
	ROOT_PATH                        = "root"
)
