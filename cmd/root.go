package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/cesc1802/auth-service/cmd/http"
	"github.com/cesc1802/auth-service/cmd/worker"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var RootCmd = &cobra.Command{
	Use:   "app",
	Short: "Authentication and Authorization service",
	Long:  "Authentication and Authorization service",
}

var configFile string

func init() {
	cobra.OnInitialize(func() {
		if configFile != "" {
			viper.SetConfigFile(configFile)
		} else {
			dir, _ := os.Getwd()
			viper.AddConfigPath(dir)
			viper.SetConfigName("config/")
		}

		replacer := strings.NewReplacer("-", "_")
		viper.SetEnvKeyReplacer(replacer)
		viper.SetConfigType("yaml")
		viper.AutomaticEnv()

		// If a config file is found, read it in.
		if err := viper.ReadInConfig(); err == nil {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		}
	})

	RootCmd.PersistentFlags().StringVar(&configFile, "config", "", "This argument is used to point to path config file")

	http.RegisterCmdRecursive(RootCmd)
	worker.RegisterWorkerCmdRecursive(RootCmd)

}

func Execute() error {
	return RootCmd.Execute()
}
