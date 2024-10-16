package initialize

import (
	"HoneyOrangeHelper/internal/global"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
)

func InitEvetnBus() {
	pubSub := gochannel.NewGoChannel(
		gochannel.Config{},
		watermill.NewStdLogger(false, false),
	)

	global.PubSub = pubSub
}
