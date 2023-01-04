package wizard

import (
	"fmt"

	"github.com/bragi-litlausson/UnstablePM/core"
	"github.com/bragi-litlausson/UnstablePM/wizard/states"
)

func StartWizard() {

	if core.ProjectFileExists() {
		fmt.Println("NOT IMPLTEMENTED")
	} else {
		initializeProject()
	}
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
