package log

import (
	"context"

	lr "github.com/sirupsen/logrus"
)

type LogrusAdapter struct {
	*lr.Logger
}

func (l *LogrusAdapter) WithContext(ctx context.Context) Logger {
	return &LogrusEntryAdapter{l.Logger.WithContext(ctx)}
}

func (l *LogrusAdapter) WithFields(data map[string]interface{}) Logger {
	return &LogrusEntryAdapter{l.Logger.WithFields(data)}
}

type LogrusEntryAdapter struct {
	*lr.Entry
}

func (e *LogrusEntryAdapter) WithFields(data map[string]interface{}) Logger {
	return &LogrusEntryAdapter{e.Entry.WithFields(data)}
}

func (e *LogrusEntryAdapter) WithContext(ctx context.Context) Logger {
	return &LogrusEntryAdapter{e.Entry.WithContext(ctx)}
}
