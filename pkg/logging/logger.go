package logging

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	core *zap.SugaredLogger
}

func NewLogger(level string) *Logger {
	var l zapcore.Level

	switch strings.ToLower(level) {
	case "error":
		l = zapcore.ErrorLevel
	case "warn":
		l = zapcore.WarnLevel
	case "info":
		l = zapcore.InfoLevel
	case "debug":
		l = zapcore.DebugLevel
	default:
		l = zapcore.InfoLevel
	}

	loggerCfg := zap.NewProductionEncoderConfig()
	loggerCfg.EncodeTime = zapcore.RFC3339TimeEncoder
	loggerCfg.EncodeLevel = zapcore.CapitalLevelEncoder
	core := zapcore.NewCore(zapcore.NewConsoleEncoder(loggerCfg), os.Stdout, l)
	logger := zap.New(core).Sugar()

	return &Logger{
		core: logger,
	}
}

func (l *Logger) Debug(args ...interface{}) {
	l.core.Debug(args)
}

func (l *Logger) Info(args ...interface{}) {
	l.core.Info(args)
}

func (l *Logger) Warn(args ...interface{}) {
	l.core.Warn(args)
}

func (l *Logger) Error(args ...interface{}) {
	l.core.Warn(args)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.core.Fatal(args)
}

func (l *Logger) Debugf(template string, args ...interface{}) {
	l.core.Debugf(template, args...)
}

func (l *Logger) Infof(template string, args ...interface{}) {
	l.core.Infof(template, args...)
}

func (l *Logger) Warnf(template string, args ...interface{}) {
	l.core.Warnf(template, args...)
}

func (l *Logger) Errorf(template string, args ...interface{}) {
	l.core.Errorf(template, args...)
}

func (l *Logger) Fatalf(template string, args ...interface{}) {
	l.core.Fatalf(template, args...)
}
