package log

import (
	"github.com/dsc-sgu/shawty/internal/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var S *zap.SugaredLogger

// Configures the global logger.
func Init() {
	encConf := zapcore.EncoderConfig{
		MessageKey:  config.C.Log.EncoderConfig.MessageKey,
		LevelKey:    config.C.Log.EncoderConfig.LevelKey,
		EncodeLevel: config.C.Log.EncoderConfig.LevelEncoder,
		TimeKey:     config.C.Log.EncoderConfig.TimeKey,
		EncodeTime:  config.C.Log.EncoderConfig.TimeEncoder,
	}

	conf := zap.Config{
		Level:            config.C.Log.Level,
		Encoding:         config.C.Log.Encoding,
		OutputPaths:      config.C.Log.OutputPaths,
		ErrorOutputPaths: config.C.Log.ErrorOutputPaths,
		Development:      config.C.Log.DevMode,
		EncoderConfig:    encConf,
	}

	S = zap.Must(conf.Build()).Sugar()
}
