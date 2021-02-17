package app

import (
	"os"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"

	"github.com/nimoism/ad-rotator/internal/log"
)

const (
	LogBackendLogrus = "logrus"
	LogBackendZap    = "zap"
)
const defaultLogBackend = LogBackendLogrus

func NewLogger(cfg LoggerConf) (log.Logger, error) {
	backendName := cfg.Backend
	if backendName == "" {
		backendName = defaultLogBackend
	}
	switch backendName {
	case LogBackendLogrus:
		return newLogrusLogger(cfg)
	case LogBackendZap:
		return newZapLogger(cfg)
	default:
		return nil, errors.Errorf("log backend '%v' is not supported", cfg.Backend)
	}
}

func newLogrusLogger(cfg LoggerConf) (*log.LogrusAdapter, error) {
	backend := logrus.New()
	if cfg.File != "" {
		logFile, err := os.OpenFile(cfg.File, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			return nil, err
		}
		backend.SetOutput(logFile)
	}
	logLevel, err := logrus.ParseLevel(cfg.Level)
	if err != nil {
		return nil, err
	}
	backend.SetLevel(logLevel)
	return &log.LogrusAdapter{Logger: backend}, nil
}

func newZapLogger(cfg LoggerConf) (*log.ZapAdapter, error) {
	zapCfg := zap.NewProductionConfig()
	if cfg.File != "" {
		zapCfg.OutputPaths = []string{cfg.File}
	}
	if err := zapCfg.Level.UnmarshalText([]byte(cfg.Level)); err != nil {
		return nil, err
	}
	backend, err := zapCfg.Build(zap.AddCallerSkip(1))
	if err != nil {
		return nil, err
	}
	return &log.ZapAdapter{Logger: backend}, nil
}
