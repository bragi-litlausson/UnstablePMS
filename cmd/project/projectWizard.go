package project

import (
	"fmt"
	"os"

	"github.com/bragi-litlausson/UnstablePM/cmd/project/states"
)

func StartWizard() {
	init := configFileExists()

	if init {
		fmt.Println("NOT IMPLTEMENTED")
	} else {
		initializeProject()
	}
}

func configFileExists() bool {
	_, err := os.Stat(".upm")
	return err == nil
}

func initializeProject() {
	projectName := states.DefineProjectName()
	fmt.Println(projectName)
	projectType := states.DefineProjectType()
	fmt.Println(projectType)
	licenseType := states.DefineLicenseType()
	fmt.Println(licenseType)
	// readme
	// shell?
}
