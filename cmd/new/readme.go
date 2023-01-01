package new

import (
	"github.com/bragi-litlausson/UnstablePM/cmd/project/states"
	"github.com/spf13/cobra"
)

var readmeCmd = &cobra.Command{
	Use:   "readme",
	Short: "Creates README.md from template",
	Long: `Create README.md from template
	In case "--name"  flag is not provided place holder name will be used.`,
	Run: func(cmd *cobra.Command, args []string) {
		states.CreateReadme(projectName)
	},
}

var projectName = ""

func init() {
	readmeCmd.Flags().StringVarP(&projectName, "name", "n", "", "Name of the project")

	NewCmd.AddCommand(readmeCmd)
}
