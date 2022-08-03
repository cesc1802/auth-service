package cmd

import (
	"fmt"
	"github.com/cesc1802/auth-service/cmd/http"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"
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
			viper.SetEnvPrefix("")
		}

		replacer := strings.NewReplacer("-", "_")
		viper.SetEnvKeyReplacer(replacer)
		viper.AutomaticEnv()

		// If a config file is found, read it in.
		if err := viper.ReadInConfig(); err == nil {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		}
	})

	RootCmd.PersistentFlags().StringVar(&configFile, "config", "", "This argument is used to point to path config file")

	http.RegisterCmdRecursive(RootCmd)

}

func Execute() error {
	return RootCmd.Execute()
}
