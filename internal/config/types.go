package config

import "go.uber.org/zap/zapcore"

type Environment string

const (
	Dev  Environment = "dev"
	Prod Environment = "prod"
)

type LoggerLevel string

const (
	TraceLevel LoggerLevel = "trace"
	DebugLevel LoggerLevel = "debug"
	InfoLevel  LoggerLevel = "info"
	WarnLevel  LoggerLevel = "warn"
	ErrorLevel LoggerLevel = "error"
	PanicLevel LoggerLevel = "panic"
)

func (ll LoggerLevel) ToZapLevel() zapcore.Level {
	switch ll {
	case TraceLevel, DebugLevel:
		return zapcore.DebugLevel
	case InfoLevel:
		return zapcore.InfoLevel
	case WarnLevel:
		return zapcore.WarnLevel
	case ErrorLevel:
		return zapcore.ErrorLevel
	case PanicLevel:
		return zapcore.PanicLevel
	default:
		return zapcore.InfoLevel
	}
}
