package logger

import (
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/config"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"
	"testing"
)

func TestNewZapLogger_checkConfiguration(t *testing.T) {
	conf := config.ApplicationConfig{
		Env: config.Prod,
		Logger: config.LoggerConfig{
			Level: config.ErrorLevel,
		},
	}
	logger, err := newZapLogger(&conf)
	assert.Nil(t, err)
	assert.Equal(t, config.ErrorLevel, logger.level)
	checkRes := logger.log.Desugar().Check(zapcore.InfoLevel, "test")
	assert.Nil(t, checkRes) //Logger won't write to log
	checkRes = logger.log.Desugar().Check(zapcore.ErrorLevel, "test")
	assert.NotNil(t, checkRes) //Logger will write to log
}

func TestNewZapLogger_notFailFromDefaultConfig(t *testing.T) {
	conf := config.ApplicationConfig{}
	logger, err := newZapLogger(&conf)
	assert.Nil(t, err)
	assert.NotNil(t, logger)
	//Check default info level
	checkRes := logger.log.Desugar().Check(zapcore.DebugLevel, "test")
	assert.Nil(t, checkRes) //Logger won't write to log
	checkRes = logger.log.Desugar().Check(zapcore.InfoLevel, "test")
	assert.NotNil(t, checkRes) //Logger will write to log
}
