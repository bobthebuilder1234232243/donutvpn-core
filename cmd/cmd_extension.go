package cmd

import (
	_ "github.com/bobthebuilder1234232243/donutvpn-core/extension/repository"
	"github.com/bobthebuilder1234232243/donutvpn-core/extension/server"
	"github.com/spf13/cobra"
)

var commandExtension = &cobra.Command{
	Use:   "extension",
	Short: "extension configuration",
	Args:  cobra.MaximumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		server.StartTestExtensionServer()
	},
}

func init() {
	// commandWarp.Flags().StringVarP(&warpKey, "key", "k", "", "warp key")
	mainCommand.AddCommand(commandExtension)
}
