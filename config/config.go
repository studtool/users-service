package config

import (
	"github.com/studtool/common/config"
)

var (
	ServerPort = config.NewStringDefault("STUDTOOL_USERS_SERVICE_PORT", "80")

	RepositoriesEnabled = config.NewFlagDefault("STUDTOOL_USERS_SERVICE_REPOSITORIES_ENABLED", false)

	StorageHost = config.NewStringDefault("STUDTOOL_USERS_STORAGE_HOST", "127.0.0.1")
	StoragePort = config.NewStringDefault("STUDTOOL_USERS_STORAGE_PORT", "27017")
	StorageDB   = config.NewStringDefault("STUDTOOL_USERS_STORAGE_NAME", "users")
)
