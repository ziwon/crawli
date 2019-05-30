package cmd

import (
	"github.com/spf13/cobra"
	"github.com/ziwon/crawli/crawli"
)

var collectCmd = &cobra.Command{
	Use:   "collect",
	Short: "Collect data with given worksheet",
	Run:   collect,
}

func init() {
	rootCmd.AddCommand(collectCmd)
}

func collect(cmd *cobra.Command, args []string) {
	crawli.Collect()
}
