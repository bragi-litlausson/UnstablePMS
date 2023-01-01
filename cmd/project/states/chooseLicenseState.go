package states

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/bragi-litlausson/UnstablePM/core"
)

func ChooseLicenseType() string {
	licenseTypes := core.GetLicenseNames()
	licensePrompt := &survey.Select{
		Message: "Choose license type",
		Options: licenseTypes,
	}

	licenseName := ""
	err := survey.AskOne(licensePrompt, &licenseName)
	if err != nil {
		panic(err)
	}

	err = core.CreateLicenseFile(licenseName)
	if err != nil {
		panic(err)
	}

	return licenseName
}
