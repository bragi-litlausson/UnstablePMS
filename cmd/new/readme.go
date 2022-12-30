package new

import (
	"os"
	"text/template"

	"github.com/spf13/cobra"
)

var readmeCmd = &cobra.Command{
	Use:   "readme",
	Short: "Creates README.md from template",
	Long: `Create README.md from template
	In case "--name"  flag is not provided place holder name will be used.`,
	Run: func(cmd *cobra.Command, args []string) {
		createReadMe()
	},
}

var projectName string

func init() {
	readmeCmd.Flags().StringVarP(&projectName, "name", "n", "", "Name of the project")

	NewCmd.AddCommand(readmeCmd)
}

func createReadMe() {
	t := template.Must(template.New("readme").Parse(readMeTemplate))
	file, err := os.Create("README.md")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if projectName == "" {
		projectName = "[name_here]"
	}

	err = t.Execute(file, projectName)
	if err != nil {
		panic(err)
	}

}

const readMeTemplate = `# {{.}}

Short description of the project

# Description

More in-depth project description

## Built with

### Library #1
Awesome library used to create this project

# Getting started

## Dependencies

Any requirements to run this

## Installing

How to install god damn thing

## Usage

Grab big hammer and...

# Known issues

FAQ it

# Contributors

Awesome people that helped

# Version history

For more detailed information check release history(Add link)

- v1.0.0
  - awesome feature
  - not so awesome feature

# License

This project is licensed under [NAME HERE] License - see LICENSE file for more details.

# Standing on shoulders of the giants

Thank you mom!`
