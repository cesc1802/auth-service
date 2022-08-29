package cache

import (
	"time"

	"github.com/cesc1802/auth-service/app_context"
	"github.com/cesc1802/auth-service/common"
	"github.com/cesc1802/auth-service/features/v1/worker/cache/transport"
	"github.com/cesc1802/auth-service/pkg/broker/rabbitmq"
	"github.com/cesc1802/auth-service/pkg/cache"
	"github.com/cesc1802/auth-service/pkg/cache/redis"
	"github.com/cesc1802/auth-service/pkg/database"
	"github.com/cesc1802/auth-service/pkg/logger"
	"github.com/cesc1802/auth-service/pkg/tokenprovider/jwt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewCacheWorkerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cache",
		Short: "New cache worker",
		Run: func(cmd *cobra.Command, args []string) {
			var l = logger.Init(
				logger.WithLogDir("logs/"),
				logger.WithDebug(true),
				logger.WithConsole(true),
			)
			defer l.Sync()

			redisClient := cache.NewRedisClient(
				viper.GetString(common.RedisHost),
				viper.GetString(common.RedisPort),
				viper.GetString(common.RedisPassword),
				viper.GetInt(common.RedisDB),
			)
			redisCache := redis.NewRedisCache(time.Hour*24, redisClient)

			gormConfig := database.NewAppGormConfig(viper.GetString(common.DbHost),
				viper.GetString(common.DbPort), viper.GetString(common.DbUserName),
				viper.GetString(common.DbName), viper.GetString(common.DbPassword))

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

			transport.LoginCacheWorker(appContext)
		},
	}
	return cmd
}

func RegisterCacheWorkerCmdRecursive(parent *cobra.Command) {
	cmd := NewCacheWorkerCommand()
	parent.AddCommand(cmd)
}
