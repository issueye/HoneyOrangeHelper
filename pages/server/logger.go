package server

import (
	"HoneyOrangeHelper/internal/global"
	"HoneyOrangeHelper/pkg/utils"
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func (f *TFrm_server) InitLogger() {
	path := fmt.Sprintf("%s/logs/%d", global.ROOT_PATH, f.data.Code)
	utils.IsExistsCreatePath(fmt.Sprintf("%s/logs", global.ROOT_PATH), fmt.Sprintf("%d", f.data.Code))

	fp := path + "/server.log"

	f.lumberJackLogger = &lumberjack.Logger{
		Filename:   fp,
		MaxSize:    100, // megabytes
		MaxBackups: 10,
		MaxAge:     28, // days
		Compress:   true,
	}

	// 自定义文件：行号输出项
	customCallerEncoder := func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString("[" + caller.TrimmedPath() + "]")
	}

	encoderConf := zapcore.EncoderConfig{
		CallerKey:      "caller_line", // 打印文件名和行数
		LevelKey:       "level_name",
		MessageKey:     "msg",
		TimeKey:        "ts",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeCaller:   customCallerEncoder, // 全路径编码器
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	writerSyncer := zapcore.AddSync(f.lumberJackLogger)
	encoder := zapcore.NewConsoleEncoder(encoderConf)

	core := zapcore.NewCore(encoder, writerSyncer, zapcore.DebugLevel)
	logger := zap.New(core)

	f.log = logger
	f.sugar = logger.Sugar()
}
