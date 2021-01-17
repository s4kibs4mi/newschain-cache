package cmd

import (
	"github.com/s4kibs4mi/newschain-cache/api"
	"github.com/spf13/cobra"
)

var (
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "serve starts the HTTP server",
		Run:   execStartServer,
	}
)

func execStartServer(cmd *cobra.Command, args []string) {
	api.StartServer()
}
