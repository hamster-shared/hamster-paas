package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var sugarLogger *zap.SugaredLogger

func InitLogger() {
	config := zap.Config{
		Encoding:         "console",
		Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:     "msg",
			LevelKey:       "level",
			TimeKey:        "ts",
			NameKey:        "logger",
			CallerKey:      "caller",
			FunctionKey:    zapcore.OmitKey,
			StacktraceKey:  zapcore.OmitKey,
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseColorLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}

	logger, err := config.Build()
	if err != nil {
		panic(err)
	}

	defer logger.Sync()
	sugarLogger = logger.WithOptions(zap.AddCallerSkip(1)).Sugar()
}

func Sugar() *zap.SugaredLogger {
	return sugarLogger
}

func Info(args ...any) {
	sugarLogger.Info(args...)
}

func Infof(template string, args ...any) {
	sugarLogger.Infof(template, args...)
}

func Debug(args ...any) {
	sugarLogger.Debug(args...)
}

func Debugf(template string, args ...any) {
	sugarLogger.Debugf(template, args...)
}

func Warn(args ...any) {
	sugarLogger.Warn(args...)
}

func Warnf(template string, args ...any) {
	sugarLogger.Warnf(template, args...)
}

func Error(args ...any) {
	sugarLogger.Error(args...)
}

func Errorf(template string, args ...any) {
	sugarLogger.Errorf(template, args...)
}
