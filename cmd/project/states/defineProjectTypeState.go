package states

import (
	"github.com/AlecAivazis/survey/v2"
)

var projectType string

// Gets list of project type names and runs prompt asking user to select one
func DefineProjectType() string {
	typeNames := getTypeNames()
	typeSelection := &survey.Select{
		Message: "Choose project type:",
		Options: typeNames,
	}
	err := survey.AskOne(typeSelection, &projectType)
	if err != nil {
		panic(err)
	}

	return projectType
}

func getTypeNames() []string {
	keys := make([]string, 0, len(ProjectTypes))
	for k, _ := range ProjectTypes {
		keys = append(keys, k)
	}

	return keys
}
