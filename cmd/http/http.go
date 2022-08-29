package http

import (
	"time"

	"github.com/cesc1802/auth-service/app_context"
	v1 "github.com/cesc1802/auth-service/cmd/http/v1"
	"github.com/cesc1802/auth-service/common"
	"github.com/cesc1802/auth-service/pkg/broker/rabbitmq"
	"github.com/cesc1802/auth-service/pkg/cache"
	"github.com/cesc1802/auth-service/pkg/cache/redis"
	"github.com/cesc1802/auth-service/pkg/database"
	"github.com/cesc1802/auth-service/pkg/httpserver"
	"github.com/cesc1802/auth-service/pkg/i18n"
	"github.com/cesc1802/auth-service/pkg/logger"
	"github.com/cesc1802/auth-service/pkg/tokenprovider/jwt"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

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
)

func registerFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().String(AccessTokenSecretKey,
		"54ccb26c0cec8f15d559fb1b9de680a3",
		"Setup the secret key for access token")
	cmd.PersistentFlags().Uint(AccessTokenExpiry, 15*60,
		"Setup expiry time for access token")
	cmd.PersistentFlags().String(RefreshTokenSecretKey, "54ccb26c0cec8f15d559fb1b9de680a3",
		"Setup the secret key for refresh token")
	cmd.PersistentFlags().Uint(RefreshTokenExpiry, 7*24*69,
		"Setup expiry time for refresh token")

	//setup env for http server
	cmd.PersistentFlags().String(ServPort, "7172", "http server port listen")
	cmd.PersistentFlags().String(ServMode, "debug", "http server mode. eg: debug/release.")

	//setup env for database
	cmd.PersistentFlags().String(DbUserName, "admin", "database username")
	cmd.PersistentFlags().String(DbPassword, "admin@1802", "database password")
	cmd.PersistentFlags().String(DbName, "auth_service", "database name")
	cmd.PersistentFlags().String(DbHost, "localhost", "database host used to connect")
	cmd.PersistentFlags().String(DbPort, "3306", "database host used to connect")
	cmd.PersistentFlags().String(DbLocation, "Local", "database location")
	cmd.PersistentFlags().Uint(DbMaxOpenConn, 20, "")
	cmd.PersistentFlags().Uint(DbMaxIdleConn, 20, "")

	cmd.PersistentFlags().String(RedisHost, "localhost", "redis host used to connect")
	cmd.PersistentFlags().String(RedisPort, "6379", "redis port used to connect")
	cmd.PersistentFlags().String(RedisPassword, "admin@1802", "redis password used to connect")
	cmd.PersistentFlags().Int(RedisDB, 0, "redis db used to connect")

	cmd.PersistentFlags().String(common.MQHost, "localhost", "rabbitmq host used to connect")
	cmd.PersistentFlags().Int(common.MQPort, 5672, "rabbitmq port used to connect")
	cmd.PersistentFlags().String(common.MQPassword, "guest", "rabbitmq password used to connect")
	cmd.PersistentFlags().String(common.MQUsername, "guest", "rabbitmq db used to connect")
	cmd.PersistentFlags().String(common.MQVhost, "/", "rabbitmq virtual host used to connect")
	cmd.PersistentFlags().String(common.MQExchangeName, "cache_exchange", "rabbitmq exchange name")
	cmd.PersistentFlags().String(common.MQExchangeType, "fanout", "rabbitmq exchange type")
	cmd.PersistentFlags().String(common.MQQueueName, "cache_queue", "rabbitmq queue name")
	cmd.PersistentFlags().String(common.MQBindingRoutingKey, "hehe", "rabbitmq binding routing key")
	cmd.PersistentFlags().String(common.MQConsumerOptionsTag, "cache_consumer", "rabbitmq consumer options tag")
	cmd.PersistentFlags().String(common.MQPublishingOptionsTag, "ProducerTagHede", "rabbitmq publishing options tag")
	cmd.PersistentFlags().String(common.MQPublishingOptionsRoutingKey, "naber", "rabbitmq publishing options routing key")

	cmd.PersistentFlags().StringSlice(ServSupportLanguages, []string{"en", "vi"},
		"server language support when response")

	_ = viper.BindPFlag(AccessTokenSecretKey, cmd.PersistentFlags().Lookup(AccessTokenSecretKey))
	_ = viper.BindPFlag(AccessTokenExpiry, cmd.PersistentFlags().Lookup(AccessTokenExpiry))
	_ = viper.BindPFlag(RefreshTokenSecretKey, cmd.PersistentFlags().Lookup(RefreshTokenSecretKey))
	_ = viper.BindPFlag(RefreshTokenExpiry, cmd.PersistentFlags().Lookup(RefreshTokenExpiry))

	_ = viper.BindPFlag(ServPort, cmd.PersistentFlags().Lookup(ServPort))
	_ = viper.BindPFlag(ServMode, cmd.PersistentFlags().Lookup(ServMode))

	_ = viper.BindPFlag(DbUserName, cmd.PersistentFlags().Lookup(DbUserName))
	_ = viper.BindPFlag(DbPassword, cmd.PersistentFlags().Lookup(DbPassword))
	_ = viper.BindPFlag(DbName, cmd.PersistentFlags().Lookup(DbName))
	_ = viper.BindPFlag(DbHost, cmd.PersistentFlags().Lookup(DbHost))
	_ = viper.BindPFlag(DbPort, cmd.PersistentFlags().Lookup(DbPort))
	_ = viper.BindPFlag(DbLocation, cmd.PersistentFlags().Lookup(DbLocation))
	_ = viper.BindPFlag(DbMaxOpenConn, cmd.PersistentFlags().Lookup(DbMaxOpenConn))
	_ = viper.BindPFlag(DbMaxIdleConn, cmd.PersistentFlags().Lookup(DbMaxIdleConn))

	_ = viper.BindPFlag(RedisHost, cmd.PersistentFlags().Lookup(RedisHost))
	_ = viper.BindPFlag(RedisPort, cmd.PersistentFlags().Lookup(RedisPort))
	_ = viper.BindPFlag(RedisPassword, cmd.PersistentFlags().Lookup(RedisPassword))
	_ = viper.BindPFlag(RedisDB, cmd.PersistentFlags().Lookup(RedisDB))

	_ = viper.BindPFlag(common.MQHost, cmd.PersistentFlags().Lookup(common.MQHost))
	_ = viper.BindPFlag(common.MQPort, cmd.PersistentFlags().Lookup(common.MQPort))
	_ = viper.BindPFlag(common.MQUsername, cmd.PersistentFlags().Lookup(common.MQUsername))
	_ = viper.BindPFlag(common.MQPassword, cmd.PersistentFlags().Lookup(common.MQPassword))
	_ = viper.BindPFlag(common.MQVhost, cmd.PersistentFlags().Lookup(common.MQVhost))
	_ = viper.BindPFlag(common.MQExchangeName, cmd.PersistentFlags().Lookup(common.MQExchangeName))
	_ = viper.BindPFlag(common.MQExchangeType, cmd.PersistentFlags().Lookup(common.MQExchangeType))
	_ = viper.BindPFlag(common.MQQueueName, cmd.PersistentFlags().Lookup(common.MQQueueName))
	_ = viper.BindPFlag(common.MQBindingRoutingKey, cmd.PersistentFlags().Lookup(common.MQBindingRoutingKey))
	_ = viper.BindPFlag(common.MQConsumerOptionsTag, cmd.PersistentFlags().Lookup(common.MQConsumerOptionsTag))
	_ = viper.BindPFlag(common.MQPublishingOptionsTag, cmd.PersistentFlags().Lookup(common.MQPublishingOptionsTag))
	_ = viper.BindPFlag(common.MQPublishingOptionsRoutingKey, cmd.PersistentFlags().Lookup(common.MQPublishingOptionsRoutingKey))

	_ = viper.BindPFlag(ServSupportLanguages, cmd.PersistentFlags().Lookup(ServSupportLanguages))
}

func NewServerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serv",
		Short: "Start HTTP Auth Service",
		Long:  "Start HTTP Auth Service",
		RunE: func(cmd *cobra.Command, args []string) error {
			var l = logger.Init(
				logger.WithLogDir("logs/"),
				logger.WithDebug(true),
				logger.WithConsole(true),
			)
			defer l.Sync()

			redisClient := cache.NewRedisClient(
				viper.GetString(RedisHost),
				viper.GetString(RedisPort),
				viper.GetString(RedisPassword),
				viper.GetInt(RedisDB),
			)
			redisCache := redis.NewRedisCache(time.Hour*24, redisClient)

			httpConfig := httpserver.NewMyHttpServerConfig(viper.GetString(ServMode),
				viper.GetString(ServPort))

			gormConfig := database.NewAppGormConfig(viper.GetString(DbHost),
				viper.GetString(DbPort), viper.GetString(DbUserName),
				viper.GetString(DbName), viper.GetString(DbPassword))

			appI18nConfig := i18n.NewI18nConfig(viper.GetStringSlice(ServSupportLanguages))
			appI18n, _ := i18n.NewI18n(appI18nConfig)

			router := httpserver.New(httpConfig, appI18n)
			appGorm := database.NewAppGorm(gormConfig)

			producer := rabbitmq.NewMQPublisher(rabbitmq.MQConfig{
				Config: rabbitmq.Config{
					Host:     viper.GetString(common.MQHost),
					Port:     viper.GetInt(common.MQPort),
					Username: viper.GetString(common.MQUsername),
					Password: viper.GetString(common.MQPassword),
					Vhost:    viper.GetString(common.MQVhost),
				},
				Exchange: rabbitmq.Exchange{
					Name: viper.GetString(common.MQExchangeName),
				},
				PublishingOptions: rabbitmq.PublishingOptions{
					Tag:        viper.GetString(common.MQPublishingOptionsTag),
					RoutingKey: viper.GetString(common.MQPublishingOptionsRoutingKey),
				},
			})
			consumer := rabbitmq.NewMQConsumer(rabbitmq.MQConfig{
				Config: rabbitmq.Config{
					Host:     viper.GetString(common.MQHost),
					Port:     viper.GetInt(common.MQPort),
					Username: viper.GetString(common.MQUsername),
					Password: viper.GetString(common.MQPassword),
					Vhost:    viper.GetString(common.MQVhost),
				},
				Exchange: rabbitmq.Exchange{
					Name:    viper.GetString(common.MQExchangeName),
					Type:    viper.GetString(common.MQExchangeType),
					Durable: true,
				},
				Queue: rabbitmq.Queue{
					Name:    viper.GetString(common.MQQueueName),
					Durable: true,
				},
				BindingOptions: rabbitmq.BindingOptions{
					RoutingKey: viper.GetString(common.MQBindingRoutingKey),
				},
				ConsumerOptions: rabbitmq.ConsumerOptions{
					Tag:     viper.GetString(common.MQConsumerOptionsTag),
					AutoAck: true,
				},
			})

			appContext := app_context.NewAppContext(appGorm.GetDB(),
				jwt.NewTokenJWTProvider(viper.GetString(common.AccessTokenExpiry), viper.GetUint(common.AccessTokenExpiry)),
				jwt.NewTokenJWTProvider(viper.GetString(common.RefreshTokenSecretKey), viper.GetUint(common.RefreshTokenExpiry)),
				redisCache,
				producer,
				consumer,
			)

			router.AddHandler(v1.SetupRoute(appContext))
			router.Start()

			gracehttp.Serve(router.Server)
			return nil
		},
	}
	return cmd
}

func RegisterCmdRecursive(parent *cobra.Command) {
	registerFlags(parent)
	cmd := NewServerCommand()
	parent.AddCommand(cmd)
}
