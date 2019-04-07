package config

import (
	"github.com/studtool/common/config"
)

var (
	ServerPort = config.NewStringDefault("STUDTOOL_USERS_SERVICE_PORT", "80")

	RepositoriesEnabled = config.NewFlagDefault("STUDTOOL_AUTH_SERVICE_REPOSITORIES_ENABLED", false)

	StorageHost     = config.NewStringDefault("STUDTOOL_USERS_STORAGE_HOST", "127.0.0.1")
	StoragePort     = config.NewStringDefault("STUDTOOL_USERS_STORAGE_PORT", "5432")
	StorageDB       = config.NewStringDefault("STUDTOOL_USERS_STORAGE_NAME", "users")
	StorageUser     = config.NewStringDefault("STUDTOOL_USERS_STORAGE_USER", "user")
	StoragePassword = config.NewStringDefault("STUDTOOL_USERS_STORAGE_PASSWORD", "password")
)
