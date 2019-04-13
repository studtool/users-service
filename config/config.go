package config

import (
	"time"

	"github.com/studtool/common/config"
)

var (
	ServerPort = config.NewStringDefault("STUDTOOL_USERS_SERVICE_PORT", "80")

	RepositoriesEnabled = config.NewFlagDefault("STUDTOOL_USERS_SERVICE_REPOSITORIES_ENABLED", false)
	QueuesEnabled       = config.NewFlagDefault("STUDTOOL_USERS_SERVICE_QUEUES_ENABLED", false)

	StorageHost = config.NewStringDefault("STUDTOOL_USERS_STORAGE_HOST", "127.0.0.1")
	StoragePort = config.NewStringDefault("STUDTOOL_USERS_STORAGE_PORT", "27017")
	StorageDB   = config.NewStringDefault("STUDTOOL_USERS_STORAGE_NAME", "users")

	UsersMqHost     = config.NewStringDefault("STUDTOOL_USERS_MQ_HOST", "127.0.0.1")
	UsersMqPort     = config.NewStringDefault("STUDTOOL_USERS_MQ_PORT", "5672")
	UsersMqUser     = config.NewStringDefault("STUDTOOL_USERS_MQ_USER", "user")
	UsersMqPassword = config.NewStringDefault("STUDTOOL_USERS_MQ_PASSWORD", "password")

	UsersMqConnNumRet = config.NewIntDefault("STUDTOOL_USERS_MQ_CONNECTION_NUM_RETRIES", 10)
	UsersMqConnRetItv = config.NewTimeSecsDefault("STUDTOOL_USERS_MQ_CONNECTION_RETRY_INTERVAL", 2*time.Second)

	CreatedUsersQueueName = config.NewStringDefault("STUDTOOL_CREATED_USERS_QUEUE_NAME", "created_users")
	DeletedUsersQueueName = config.NewStringDefault("STUDTOOL_DELETED_USERS_QUEUE_NAME", "deleted_users")
)
