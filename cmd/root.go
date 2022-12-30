package cmd

import (
	"os"

	"github.com/bragi-litlausson/UnstablePM/cmd/new"
	"github.com/bragi-litlausson/UnstablePM/cmd/project"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "upm",
	Short: "Project manager stuck in gophers body",
	Long:  ``,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(new.NewCmd)
	rootCmd.AddCommand(project.ProjectCmd)
}
