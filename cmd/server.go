package cmd

import "github.com/spf13/cobra"

var (
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "serve starts the HTTP server",
		Run:   execStartServer,
	}
)

func execStartServer(cmd *cobra.Command, args []string) {

}
