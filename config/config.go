package config

import (
	"time"

	"github.com/studtool/common/config"
	"github.com/studtool/common/logs"
)

var (
	// TODO compile-time injection
	// nolint:golint,gochecknoglobals
	ComponentName = "users-service"

	// TODO compile-time injection
	// nolint:golint,gochecknoglobals
	ComponentVersion = "v0.0.1"

	_ = func() *config.FlagVar {
		f := config.NewFlagDefault("STUDTOOL_USERS_SERVICE_SHOULD_LOG_ENV_VARS", false)
		if f.Value() {
			config.SetLogger(logs.NewRawLogger())
		}
		return f
	}()

	ServerPort = config.NewIntDefault("STUDTOOL_USERS_SERVICE_PORT", 80)

	// nolint:golint,gochecknoglobals
	CorsAllowed = config.NewFlagDefault("STUDTOOL_USERS_SERVICE_SHOULD_ALLOW_CORS", false)

	// nolint:golint,gochecknoglobals
	RequestsLogsEnabled = config.NewFlagDefault("STUDTOOL_USERS_SERVICE_SHOULD_LOG_REQUESTS", true)

	RepositoriesEnabled = config.NewFlagDefault("STUDTOOL_USERS_SERVICE_REPOSITORIES_ENABLED", false)
	QueuesEnabled       = config.NewFlagDefault("STUDTOOL_USERS_SERVICE_QUEUES_ENABLED", false)

	StorageHost = config.NewStringDefault("STUDTOOL_USERS_STORAGE_HOST", "127.0.0.1")
	StoragePort = config.NewIntDefault("STUDTOOL_USERS_STORAGE_PORT", 27017)
	StorageDB   = config.NewStringDefault("STUDTOOL_USERS_STORAGE_NAME", "users")

	MqHost     = config.NewStringDefault("STUDTOOL_MQ_HOST", "127.0.0.1")
	MqPort     = config.NewIntDefault("STUDTOOL_MQ_PORT", 5672)
	MqUser     = config.NewStringDefault("STUDTOOL_MQ_USER", "user")
	MqPassword = config.NewStringDefault("STUDTOOL_MQ_PASSWORD", "password")

	MqConnNumRet = config.NewIntDefault("STUDTOOL_MQ_CONNECTION_NUM_RETRIES", 10)
	MqConnRetItv = config.NewTimeDefault("STUDTOOL_MQ_CONNECTION_RETRY_INTERVAL", 2*time.Second)
)
