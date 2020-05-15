package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var anywhereCommands = []string{"version", "info", "help"}

var RootCmd = &cobra.Command{
	SilenceErrors: true,
	Use:           "网关服务",
	Short:         "网关服务",
	PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
		isFreeCommand := false
		for _, freeCmd := range anywhereCommands {
			if freeCmd == cmd.Name() {
				isFreeCommand = true
			}
		}
		if isFreeCommand {
			return nil
		}
		return nil
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
