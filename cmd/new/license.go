package new

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// addLicenseCmd represents the addLicense command
var licenseCmd = &cobra.Command{
	Use:   "license",
	Short: "Adds 0-BSD license to the project",
	Long:  `Adds 0-BSD license to the project.`,
	Run: func(cmd *cobra.Command, args []string) {
		createZeroLicense()
	},
}

func init() {
	NewCmd.AddCommand(licenseCmd)
}

func createZeroLicense() {
	const zeroBSD = "Permission to use, copy, modify, and/or distribute this software for any purpose with or without fee is hereby granted.\nTHE SOFTWARE IS PROVIDED \"AS IS\" AND THE AUTHOR DISCLAIMS ALL WARRANTIES WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE."

	f, err := os.Create("LICENSE")
	if err != nil {
		panic(err)
	}

	_, err = f.WriteString(zeroBSD)
	if err != nil {
		panic(err)
	}

	f.Close()

	fmt.Println("Created 0-BSD LICENSE")
}
