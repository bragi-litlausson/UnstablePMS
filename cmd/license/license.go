package license

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// addLicenseCmd represents the addLicense command
var CMD = &cobra.Command{
	Use:   "addLicense",
	Short: "Adds 0-BSD license to the project",
	Long:  `Adds 0-BSD license to the project.`,
	Run: func(cmd *cobra.Command, args []string) {
		createZeroLicense()
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addLicenseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addLicenseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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
