package pg

import (
	"context"

	"github.com/jackc/pgx/v4"

	"github.com/nimoism/ad-rotator/internal/log"
)

type PgxLoggerAdapter struct {
	log.Logger
}

func (p *PgxLoggerAdapter) Log(ctx context.Context, level pgx.LogLevel, msg string, data map[string]interface{}) {
	logger := p.Logger.WithContext(ctx)
	if data != nil {
		logger = logger.WithFields(data)
	}
	var logFunc func(...interface{})
	switch level {
	case pgx.LogLevelNone:
		return
	case pgx.LogLevelTrace:
		logFunc = logger.Trace
	case pgx.LogLevelDebug:
		logFunc = logger.Debug
	case pgx.LogLevelInfo:
		logFunc = logger.Info
	case pgx.LogLevelWarn:
		logFunc = logger.Warn
	case pgx.LogLevelError:
		logFunc = logger.Error
	default:
		logger.Errorf("Unknown log level: %v", level)
		logFunc = logger.Error
	}
	logFunc(msg)
}
