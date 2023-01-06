package states

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/bragi-litlausson/UnstablePMS/core"
)

func RunProjectTypeState() core.ProjectType {
	typeNames := getTypeNames()
	typeSelection := &survey.Select{
		Message: "Choose project type:",
		Options: typeNames,
	}

	projectType := ""
	err := survey.AskOne(typeSelection, &projectType)
	if err != nil {
		panic(err)
	}

	return core.ProjectTypes[projectType]
}

func getTypeNames() []string {
	return core.GetMapKeys(core.ProjectTypes)
}
