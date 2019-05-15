package config

import (
	"time"

	"github.com/studtool/common/config"

	"github.com/studtool/users-service/beans"
)

var (
	_ = func() *cconfig.FlagVar {
		f := cconfig.NewFlagDefault("STUDTOOL_USERS_SERVICE_SHOULD_LOG_ENV_VARS", false)
		if f.Value() {
			cconfig.SetLogger(beans.Logger)
		}
		return f
	}()

	ServerPort = cconfig.NewStringDefault("STUDTOOL_USERS_SERVICE_PORT", "80")

	ShouldLogRequests = cconfig.NewFlagDefault("STUDTOOL_USERS_SERVICE_SHOULD_LOG_REQUEST", true)

	RepositoriesEnabled = cconfig.NewFlagDefault("STUDTOOL_USERS_SERVICE_REPOSITORIES_ENABLED", false)
	QueuesEnabled       = cconfig.NewFlagDefault("STUDTOOL_USERS_SERVICE_QUEUES_ENABLED", false)

	StorageHost = cconfig.NewStringDefault("STUDTOOL_USERS_STORAGE_HOST", "127.0.0.1")
	StoragePort = cconfig.NewIntDefault("STUDTOOL_USERS_STORAGE_PORT", 27017)
	StorageDB   = cconfig.NewStringDefault("STUDTOOL_USERS_STORAGE_NAME", "users")

	UsersMqHost     = cconfig.NewStringDefault("STUDTOOL_USERS_MQ_HOST", "127.0.0.1")
	UsersMqPort     = cconfig.NewIntDefault("STUDTOOL_USERS_MQ_PORT", 5672)
	UsersMqUser     = cconfig.NewStringDefault("STUDTOOL_USERS_MQ_USER", "user")
	UsersMqPassword = cconfig.NewStringDefault("STUDTOOL_USERS_MQ_PASSWORD", "password")

	UsersMqConnNumRet = cconfig.NewIntDefault("STUDTOOL_USERS_MQ_CONNECTION_NUM_RETRIES", 10)
	UsersMqConnRetItv = cconfig.NewTimeDefault("STUDTOOL_USERS_MQ_CONNECTION_RETRY_INTERVAL", 2*time.Second)

	CreatedUsersQueueName = cconfig.NewStringDefault("STUDTOOL_CREATED_USERS_QUEUE_NAME", "created_users")
	DeletedUsersQueueName = cconfig.NewStringDefault("STUDTOOL_DELETED_USERS_QUEUE_NAME", "deleted_users")
)
