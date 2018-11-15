package cmd

import (
	"github.com/filkra/attendance/cmd/list"
	"github.com/spf13/cobra"
	"os"
)

var rootCommand = &cobra.Command{
	Use:   "attendance",
	Short: "A simple command line tool for attendance list generation",
	SilenceErrors: true,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(1)
	},
}

func init() {
	rootCommand.AddCommand(list.RootCommand)
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		os.Exit(1)
	}
}