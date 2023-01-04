package wizard

import (
	"fmt"
	"os"

	"github.com/bragi-litlausson/UnstablePM/core"
	"github.com/bragi-litlausson/UnstablePM/wizard/states"
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
	data := new(core.ProjectData)
	data.ProjectName = states.RunProjectNameState()
	data.ProjectType = states.RunProjectTypeState()
	data.LicenseType = states.RunLicenseState()
	states.RunReadmeState(data.ProjectName)
	states.RunNixShellState(data)

	states.RunProjectFilesState(data)
}
