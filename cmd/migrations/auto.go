package migrations

import (
	"context"
	"github.com/s4kibs4mi/newschain-cache/app"
	"github.com/s4kibs4mi/newschain-cache/log"
	"github.com/spf13/cobra"
)

var (
	MigrationAutoCmd = &cobra.Command{
		Use:   "auto",
		Short: "auto creates the migration table",
		Run:   migrationAuto,
	}
)

func migrationAuto(cmd *cobra.Command, args []string) {
	if err := app.DB().Schema.Create(context.Background()); err != nil {
		log.Log().Errorln(err)
	}
}
