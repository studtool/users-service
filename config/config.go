package config

import (
	"github.com/studtool/common/config"
)

var (
	ServerPort = config.NewStringDefault("STUDTOOL_USERS_SERVICE_PORT", "80")
)
