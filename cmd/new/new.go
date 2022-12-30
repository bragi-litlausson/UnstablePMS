package new

import (
	"github.com/spf13/cobra"
)

var NewCmd = &cobra.Command{
	Use:   "new",
	Short: "Create essential repo files",
	Long:  "Create essential repo files",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
}
