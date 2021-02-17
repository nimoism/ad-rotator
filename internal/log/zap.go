package log

import (
	"context"
	"strconv"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapAdapter struct {
	*zap.Logger
}

func (l *ZapAdapter) log(levelFunc func(msg string, fields ...zap.Field), args ...interface{}) {
	if len(args) < 1 {
		l.Logger.Error("no message for log")
		return
	}
	var msg string
	switch m := args[0].(type) {
	case string:
		msg = m
	case error:
		msg = m.Error()
	default:
		l.Logger.Error("log message should be a string")
	}
	if len(args) > 1 {
		l.Logger.Warn("FUCK!")
	}
	levelFunc(msg)
}

func (l *ZapAdapter) logf(levelFunc func(msg string, fields ...zap.Field), msg string, args ...interface{}) {
	fields := make([]zapcore.Field, 0, len(args))
	for i, val := range args {
		fields = append(fields, zap.Reflect(strconv.Itoa(i), val))
	}
	levelFunc(msg, fields...)
}

func (l *ZapAdapter) Trace(args ...interface{}) { l.log(l.Logger.Debug, args...) }
func (l *ZapAdapter) Debug(args ...interface{}) { l.log(l.Logger.Debug, args...) }
func (l *ZapAdapter) Info(args ...interface{})  { l.log(l.Logger.Info, args...) }
func (l *ZapAdapter) Warn(args ...interface{})  { l.log(l.Logger.Warn, args...) }
func (l *ZapAdapter) Error(args ...interface{}) { l.log(l.Logger.Error, args...) }

func (l *ZapAdapter) Tracef(msg string, args ...interface{}) { l.logf(l.Logger.Debug, msg, args...) }
func (l *ZapAdapter) Debugf(msg string, args ...interface{}) { l.logf(l.Logger.Debug, msg, args...) }
func (l *ZapAdapter) Infof(msg string, args ...interface{})  { l.logf(l.Logger.Info, msg, args...) }
func (l *ZapAdapter) Warnf(msg string, args ...interface{})  { l.logf(l.Logger.Warn, msg, args...) }
func (l *ZapAdapter) Errorf(msg string, args ...interface{}) { l.logf(l.Logger.Error, msg, args...) }

func (l *ZapAdapter) WithContext(_ context.Context) Logger {
	return l
}

func (l *ZapAdapter) WithFields(data map[string]interface{}) Logger {
	fields := make([]zap.Field, 0, len(data))
	for key, val := range data {
		fields = append(fields, zap.Reflect(key, val))
	}
	return &ZapAdapter{Logger: l.Logger.With(fields...)}
}
