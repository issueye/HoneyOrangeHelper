package helper_cmd

type appLogWriter struct {
	msgChannel chan string
}

func (w *appLogWriter) Write(p []byte) (n int, err error) {
	w.msgChannel <- string(p)
	// msg := message.NewMessage(watermill.NewUUID(), p)
	// global.PubSub.Publish(global.TOPIC_CONSOLE_LOG, msg)
	return len(p), nil
}
