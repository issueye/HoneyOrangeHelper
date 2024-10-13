package helper_cmd

import "HoneyOrangeHelper/internal/global"

type appLogWriter struct{}

func (appLogWriter) Write(p []byte) (n int, err error) {
	global.MsgChannel <- string(p)
	// msg := message.NewMessage(watermill.NewUUID(), p)
	// global.PubSub.Publish(global.TOPIC_CONSOLE_LOG, msg)
	return len(p), nil
}
