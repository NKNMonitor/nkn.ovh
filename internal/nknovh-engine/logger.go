package nknovh_engine

import (
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger
var level zap.AtomicLevel
var once sync.Once

func InitLogger(logLevel string) {
	switch logLevel {
	case "DEBUG":
		level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	case "ERROR":
		level = zap.NewAtomicLevelAt(zapcore.ErrorLevel)
	case "WARN":
		level = zap.NewAtomicLevelAt(zapcore.WarnLevel)
	default:
		level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	}
	conf := zap.Config{
		Encoding:    "json",
		Level:       level,
		OutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "message",
			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
			LevelKey:     "level",
			EncodeLevel:  zapcore.CapitalLevelEncoder,

			TimeKey:    "time",
			EncodeTime: zapcore.ISO8601TimeEncoder,
		},
	}
	Logger, _ = conf.Build()
	zap.ReplaceGlobals(Logger)

}

func GetLogger(logLevel string) *zap.Logger {
	once.Do(func() { InitLogger(logLevel) })
	return Logger
}
