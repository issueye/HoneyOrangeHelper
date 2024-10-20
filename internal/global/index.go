package global

import (
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"go.uber.org/zap"
)

var (
	MsgChannel = make(chan string, 10)
	PubSub     *gochannel.GoChannel
	Sugared    *zap.SugaredLogger
	Logger     *zap.Logger
)

const (
	TOPIC_CONSOLE_LOG   = "TOPIC_CONSOLE_LOG"
	TOPIC_SERVER_REMOVE = "TOPIC_SERVER_REMOVE"
	TOPIC_SERVER_INDEX  = "TOPIC_SERVER_INDEX"
	ROOT_PATH           = "root"
)
