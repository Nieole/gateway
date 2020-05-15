package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.SetHelpCommand(helpCmd)
}

var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "帮助信息",
	Run: func(c *cobra.Command, args []string) {
		logrus.Info("获取帮助信息")
	},
}
