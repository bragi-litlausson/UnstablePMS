package states

import "github.com/AlecAivazis/survey/v2"

func RunProjectNameState() string {
	namePrompt := &survey.Input{
		Message: "Name of the project:",
	}

	name := ""
	err := survey.AskOne(namePrompt, &name)
	if err != nil {
		panic(err)
	}

	return name
}
