package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Info(message string, fields ...zap.Field)
	Warn(message string, fields ...zap.Field)
	Error(message string, fields ...zap.Field)
	Fatal(message string, fields ...zap.Field)
	Debug(message string, fields ...zap.Field)
}

// LogHandler implements Logger
type LogHandler struct {
	log *zap.Logger
	// appName    string
	// appVersion string
}

func (lh *LogHandler) Info(message string, fields ...zap.Field) {
	lh.Info(message, fields...)
}

func (lh *LogHandler) Warn(message string, fields ...zap.Field) {
	lh.Warn(message, fields...)
}

func (lh *LogHandler) Fatal(message string, fields ...zap.Field) {
	lh.Fatal(message, fields...)
}

func (lh *LogHandler) Debug(message string, fields ...zap.Field) {
	lh.Debug(message, fields...)
}

func (lh *LogHandler) Error(message string, fields ...zap.Field) {
	lh.Error(message, fields...)
}

// New returns a new instance of LogHandler
func New(skipStack uint) *LogHandler {
	config := zap.NewProductionConfig()
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.StacktraceKey = ""
	config.EncoderConfig = encoderConfig

	logger, err := config.Build(zap.AddCallerSkip(int(skipStack)))
	if err != nil {
		panic(err)
	}

	return &LogHandler{
		log: logger,
	}
}
