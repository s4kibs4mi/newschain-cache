package cmd

import (
	"github.com/s4kibs4mi/newschain-cache/config"
	"github.com/s4kibs4mi/newschain-cache/log"
	"github.com/spf13/cobra"
	"os"
)

var (
	RootCmd = &cobra.Command{
		Use:   "newschain",
		Short: "A HTTP RESTful API",
	}
)

func init() {
	RootCmd.AddCommand(serveCmd)
}

func Execute() {
	if err := config.LoadConfig(); err != nil {
		log.Log().Errorln(err)
		os.Exit(1)
	}

	if err := RootCmd.Execute(); err != nil {
		log.Log().Errorln(err)
		os.Exit(1)
	}
}
