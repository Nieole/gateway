package cmd

import (
	"gateway/config"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "打印程序版本",
	Run: func(c *cobra.Command, args []string) {
		logrus.Infof("网关程序版本为: %s", config.Version)
	},
}
