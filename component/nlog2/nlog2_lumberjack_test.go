package nlog2

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"lu-short/common/run_env"
	"testing"
)

// zap log

func TestZapLog2(t *testing.T) {

	logFile := &lumberjack.Logger{
		Filename:   fmt.Sprintf("log-%s.log", run_env.GetHostName()), //filePath
		MaxSize:    1,                                                // megabytes
		MaxBackups: 4,
		MaxAge:     1,     //days
		Compress:   false, // disabled by default
	}
	defer logFile.Close()

	/*zap 的 Config 非常的繁琐也非常强大，可以控制打印 log 的所有细节，因此对于我们开发者是友好的，有利于二次封装。
	  但是对于初学者则是噩梦。因此 zap 提供了一整套的易用配置，大部分的姿势都可以通过一句代码生成需要的配置。
	*/
	enConfig := zap.NewProductionEncoderConfig() //生成配置

	// 时间格式
	enConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(enConfig), //编码器配置
		zapcore.AddSync(logFile),         //打印到控制台和文件
		zap.InfoLevel,                    //日志等级
	)

	log := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	for i := 0; i < 1000000; i++ {
		log.Info("test for test")
	}
}
