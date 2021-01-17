package cmd

import (
	"github.com/s4kibs4mi/newschain-cache/cmd/migrations"
	"github.com/spf13/cobra"
)

var (
	migrationCmd = &cobra.Command{
		Use:   "migration",
		Short: "migration executes database migrations",
	}
)

func init() {
	migrationCmd.AddCommand(migrations.MigrationAutoCmd)
}
