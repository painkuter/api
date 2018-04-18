package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitConfig(t *testing.T) {
	cfg := InitConfig("../../config/config.sample.toml")
	checkConfig := &Config{

		IsDebug: true,

		LogDir:    "log",
		LogPrefix: "api",
	}
	assert.Equal(t, cfg, checkConfig)
}
