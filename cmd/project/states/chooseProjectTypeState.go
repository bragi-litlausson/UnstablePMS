package states

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/bragi-litlausson/UnstablePM/core"
)

func ChooseProjectType() string {
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

	return projectType
}

func getTypeNames() []string {
	return core.GetMapKeys(ProjectTypes)
}
