package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version",
	Long:  "Version string",
	Run: func(cmd *cobra.Command, args []string) {
		println(fmt.Sprintf("Version: %v", viper.GetString("Version")))
	},
}

func init() {
	viper.SetDefault("Version", "0.0.0")
	rootCmd.AddCommand(versionCmd)
}
