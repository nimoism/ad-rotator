package app

import (
	"time"
)

type Config struct {
	Logger  LoggerConf
	Storage StorageConf
	API     NetConf
	Stream  StreamConf
}

type NetConf struct {
	Listen string
}

type LoggerConf struct {
	Backend string
	File    string
	Level   string
}

type StorageConf struct {
	DSN           string
	ConnTimeout   time.Duration
	MaxIdleConn   int `mapstructure:"max_idle_conn"`
	MaxConn       int `mapstructure:"max_conn"`
	Retry         int
	RetryInterval time.Duration `mapstructure:"retry_interval"`
}

type StreamConf struct {
	Host string
}
