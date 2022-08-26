package worker

import (
	"github.com/cesc1802/auth-service/cmd/worker/cache"
	"github.com/spf13/cobra"
)

func init() {
	cache.RegisterCacheWorkerCmdRecursive(workerCmd)
}

var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "Start new worker",
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			panic(err)
		}
	},
}

func RegisterWorkerCmdRecursive(parent *cobra.Command) {
	parent.AddCommand(workerCmd)
}
