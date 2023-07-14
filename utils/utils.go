package utils

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type PrefixedFormatter struct {
	Prefix string
	logrus.Formatter
	TextColor color.Attribute
}

func (f *PrefixedFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	prefix := color.New(f.TextColor).SprintFunc()(f.Prefix)
	entry.Message = fmt.Sprintf("%s %s", prefix, entry.Message)
	return f.Formatter.Format(entry)
}

const (
	Debug = zap.DebugLevel
	Info  = zap.InfoLevel
	Warn  = zap.WarnLevel
	Error = zap.ErrorLevel
	Fatal = zap.FatalLevel
)

func NewZapLogger(component string, level zapcore.Level) *zap.Logger {
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:    "message",
		LevelKey:      "level",
		TimeKey:       "time",
		NameKey:       "logger",
		CallerKey:     "caller",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.LowercaseLevelEncoder,
		//EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	core := zapcore.NewCore(
		//zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.Lock(os.Stdout),
		level,
	)

	prefixedLogger := zap.New(core, zap.Fields(zap.String("component", component)), zap.AddCaller())

	return prefixedLogger
}
