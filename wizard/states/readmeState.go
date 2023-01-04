package states

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/bragi-litlausson/UnstablePM/core"
)

func RunReadmeState(projectName string) {
	prompt := &survey.Select{
		Message: "Create README.md",
		Options: []string{core.TemplateReadme, core.EmptyReadme, core.NoReadme},
	}

	choice := ""
	survey.AskOne(prompt, &choice)
	core.CreateReadmeFile(choice, projectName)
}
