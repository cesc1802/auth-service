package migration

import (
	"database/sql"
	"fmt"
	"os"
	"path"

	"github.com/cesc1802/auth-service/common"
	"github.com/cesc1802/auth-service/pkg/database"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewMigrationCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "migrate",
		Short: "migrate to db",
		Run: func(cmd *cobra.Command, args []string) {
			gormConfig := database.NewAppGormConfig(viper.GetString(common.DbHost),
				viper.GetString(common.DbPort), viper.GetString(common.DbUserName),
				viper.GetString(common.DbName), viper.GetString(common.DbPassword))

			db, err := sql.Open("mysql", gormConfig.Uri())
			if err != nil {
				panic(err)
			}

			rootPath, _ := os.Getwd()
			migration := &migrate.FileMigrationSource{
				Dir: path.Join(rootPath, "migrations"),
			}

			n, err := migrate.Exec(db, "mysql", migration, migrate.Up)
			if err != nil {
				panic(err)
			}

			fmt.Printf("Applied %d migrations!\n", n)
		},
	}

	return cmd
}

func RegisterCmdRecursive(parent *cobra.Command) {
	cmd := NewMigrationCommand()
	parent.AddCommand(cmd)
}
