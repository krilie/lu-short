package nlog

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"lu-short/common/run_env"
	"lu-short/component/ncfg"
	"os"
)

type NLog struct {
	log      *zap.Logger
	LogFiles io.Closer
}

func (n *NLog) Get(ctx context.Context) *zap.Logger {
	return n.log
}

func (n *NLog) GetSugar(ctx context.Context) *zap.SugaredLogger {
	return n.log.Sugar()
}

func (n *NLog) Close() {
	err := n.log.Sync()
	if err != nil {
		panic(err)
	}
	err = n.LogFiles.Close()
	if err != nil {
		panic(err)
	}
}

func NewNLog(cfg *ncfg.NConfig) *NLog {
	var logFile io.WriteCloser
	if cfg.GetLogCfg().LogFile == "stdout" {
		logFile = os.Stdout
	} else {
		logFile = &lumberjack.Logger{
			Filename:   fmt.Sprintf("log/log-%s.log", run_env.GetHostName()), //filePath
			MaxSize:    100,                                                  // megabytes
			MaxBackups: 7,
			MaxAge:     30,    //days
			Compress:   false, // disabled by default
		}
	}
	var level zapcore.Level
	err := level.UnmarshalText([]byte(cfg.GetLogCfg().LogLevel)) //日志等级
	if err != nil {
		panic(err)
	}

	enConfig := zap.NewProductionEncoderConfig() //生成配置
	// 时间格式
	enConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(enConfig), //编码器配置
		zapcore.AddSync(logFile),         //打印到控制台和文件
		level,                            //日志等级
	)

	log := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	return &NLog{log: log, LogFiles: logFile}
}
