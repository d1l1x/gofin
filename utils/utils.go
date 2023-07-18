package utils

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type Stack []interface{}

func (s *Stack) Push(v interface{}) {
	*s = append(*s, v)
}

func (s *Stack) Pop() interface{} {
	// Get the last value of the stack
	res := (*s)[len(*s)-1]
	// Remove the last value of the stack
	*s = (*s)[:len(*s)-1]
	return res
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
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

var log = NewZapLogger("utils", Debug)
