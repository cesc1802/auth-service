package common

import "time"

const (
	RequesterKey = "requester"

	UserRoleCacheKey       = "user:%v:roles"
	UserPermissionCacheKey = "user:%v:permissions"

	DefaultCacheExpiration = time.Hour * 24
)

// command constants
const (
	DbUserName    = "db-username"
	DbPassword    = "db-password"
	DbName        = "db-name"
	DbHost        = "db-host"
	DbPort        = "db-port"
	DbCharset     = "db-charset"
	DbLocation    = "db-location"
	DbMaxOpenConn = "db-max-open-conn"
	DbMaxIdleConn = "db-max-idle-conn"

	AccessTokenSecretKey = "access-secret-key"
	AccessTokenExpiry    = "access-expiry"

	RefreshTokenSecretKey = "refresh-secret-key"
	RefreshTokenExpiry    = "refresh-expiry"

	ServPort             = "server-port"
	ServMode             = "server-mode"
	ServSupportLanguages = "server-support-languages"

	RedisHost     = "redis-host"
	RedisPort     = "redis-port"
	RedisPassword = "redis-password"
	RedisDB       = "redis-db"

	MQHost                        = "mq-host"
	MQPort                        = "mq-port"
	MQPassword                    = "mq-password"
	MQUsername                    = "mq-username"
	MQVhost                       = "mq-vhost"
	MQExchangeName                = "mq-exchange-name"
	MQExchangeType                = "mq-exchange-type"
	MQQueueName                   = "mq-queue-name"
	MQBindingRoutingKey           = "mq-binding-routing-key"
	MQConsumerOptionsTag          = "mq-consumer-option-tag"
	MQPublishingOptionsTag        = "mq-publishing-option-tag"
	MQPublishingOptionsRoutingKey = "mq-publishing-option-routing-key"
)
