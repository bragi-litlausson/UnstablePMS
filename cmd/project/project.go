package project

import (
	"github.com/bragi-litlausson/UnstablePM/wizard"
	"github.com/spf13/cobra"
)

var ProjectCmd = &cobra.Command{
	Use:   "project",
	Short: "Manage project",
	Long: `Allows to set up or edit project 
`,
	Run: func(cmd *cobra.Command, args []string) {
		wizard.StartWizard()
	},
}

func init() {

}
